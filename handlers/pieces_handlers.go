package handlers

import (
	"backend/db"
	"backend/models"
	"backend/utils"
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

type gpx struct {
	Name string `json:"name"`
}

type pieceInTrack struct {
	Name string  `json:"name"`
	Time float64 `json:"time"`
}

func getAddedPieces(c echo.Context) error {
	username := c.Get("username")
	if username == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "you are not authorized",
		})
	}
	database := db.Database{}.GetConnection()
	var files []models.Piece
	rows, err := database.Query("SELECT name, filename from pieces JOIN subscribed_pieces ON subscribed_pieces.piece_id = pieces.id WHERE user_id = (SELECT id FROM users WHERE username = $1)", username)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var file models.Piece
		if err := rows.Scan(&file.Name, &file.Filename); err != nil {
			panic(err)
		}
		files = append(files, file)
	}
	return c.JSON(http.StatusOK, files)
}

func calculatePieces(c echo.Context) error {
	username := c.Get("username")
	if username == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "you are not authorized",
		})
	}
	decoder := json.NewDecoder(c.Request().Body)
	var trackName gpx
	err := decoder.Decode(&trackName)
	if err != nil {
		panic(err)
	}
	usernameStr, ok := username.(string)
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "username is not a string",
		})
	}
	trackPath := filepath.Join("resources", usernameStr, trackName.Name)
	trackFile, err := utils.ParseGPX(trackPath)
	if err != nil {
		panic(err)
	}
	trackPoints := utils.GetTrackPointsWithTime(trackFile)
	res := []pieceInTrack{}
	database := db.Database{}.GetConnection()
	rows, err := database.Query("SELECT name, filename from pieces JOIN subscribed_pieces ON subscribed_pieces.piece_id = pieces.id WHERE user_id = (SELECT id FROM users WHERE username = $1)", username)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var piece models.Piece
		if err := rows.Scan(&piece.Name, &piece.Filename); err != nil {
			panic(err)
		}
		pieceFilename := filepath.Join("static", piece.Filename)
		pieceFile, err := utils.ParseGPX(pieceFilename)
		if err != nil {
			panic(err)
		}
		piecePoints := utils.GetTrackPoints(pieceFile)
		contains, duration := utils.ContainsRouteSlidingWindow(trackPoints, piecePoints, 20.0)
		if contains {
			res = append(res, pieceInTrack{Name: piece.Name, Time: duration.Seconds()})
		}
	}
	return c.JSON(http.StatusOK, res)

}
