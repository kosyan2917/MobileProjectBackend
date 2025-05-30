package handlers

import (
	"backend/db"
	"backend/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

type trackBody struct {
	Files []track
}

type track struct {
	Name string
}

func listFiles(c echo.Context) error {
	username := c.Get("username")
	if username == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "you are not authorized",
		})
	}
	entries, _ := os.ReadDir(fmt.Sprintf("resources/%s", username))
	result := []string{}
	for _, entry := range entries {
		result = append(result, entry.Name())
	}

	return c.JSON(http.StatusOK, map[string][]string{
		"files": result,
	})
}

func getDiff(c echo.Context) error {
	decoder := json.NewDecoder(c.Request().Body)
	var t trackBody
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
	username := c.Get("username")
	if username == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "you are not authorized",
		})
	}
	database := db.Database{}.GetConnection()
	var files []models.Tracks
	rows, err := database.Query("SELECT tracks.name, time, created_at, distance from tracks INNER JOIN users ON tracks.owner_id = users.id WHERE users.username = $1", username)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var file models.Tracks
		if err := rows.Scan(&file.Name, &file.Time, &file.Created_at, &file.Distance); err != nil {
			panic(err)
		}
		files = append(files, file)
	}
	result := []models.Tracks{}
	for _, serverTrack := range files {
		flag := false
		for _, clientTrack := range t.Files {
			if serverTrack.Name == clientTrack.Name {
				flag = true
				break
			}
		}
		if !flag {
			result = append(result, serverTrack)
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"files": result,
	})

}

func getFile(c echo.Context) error {
	username := c.Get("username")
	if username == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "you are not authorized",
		})
	}
	file := c.Param("file")
	fmt.Println(file)
	decodedFile, _ := url.QueryUnescape(file)
	return c.File(fmt.Sprintf("resources/%s/%s", username, decodedFile))
}

func uploadTrack(c echo.Context) error {
	username := c.Get("username")
	if username == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "you are not authorized",
		})
	}
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.String(http.StatusBadRequest, "field ‘file’ is required")
	}
	src, err := fileHeader.Open()
	if err != nil {
		return c.String(http.StatusInternalServerError, "could not open uploaded file")
	}
	defer src.Close()
	uploadDir := filepath.Join("resources", username.(string))
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return c.String(http.StatusInternalServerError, "could not create upload directory")
	}

	// 4. Создаем файл на диске с тем же именем, что пришло от клиента
	dstPath := filepath.Join(uploadDir, fileHeader.Filename+".gpx")
	dst, err := os.Create(dstPath)
	if err != nil {
		return c.String(http.StatusInternalServerError, "could not create file on disk")
	}
	defer dst.Close()

	// 5. Копируем данные из multipart в файл
	if _, err := io.Copy(dst, src); err != nil {
		return c.String(http.StatusInternalServerError, "could not save file")
	}

	// 6. Возвращаем подтверждение
	return c.String(http.StatusCreated, "uploaded: "+fileHeader.Filename)
}
