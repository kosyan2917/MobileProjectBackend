package models

type Tracks struct {
	Name       string  `json:"name"`
	Time       int     `json:"time"`
	Created_at int     `json:"createdAt"`
	Distance   float64 `json:"distance"`
}

type Piece struct {
	Name     string `json:"name"`
	Filename string `json:"filename"`
}
