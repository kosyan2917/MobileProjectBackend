package utils

import (
	"encoding/xml"
	"io"
	"os"
	"time"
)

// GPX represents the root element of a GPX file.
type GPX struct {
	XMLName xml.Name `xml:"gpx"`
	Version string   `xml:"version,attr"`
	Creator string   `xml:"creator,attr"`
	Tracks  []Trk    `xml:"trk"`
}

// Trk represents a track.
type Trk struct {
	Segments []Trkseg `xml:"trkseg"`
}

// Trkseg represents a track segment.
type Trkseg struct {
	Points []Trkpt `xml:"trkpt"`
}

// Trkpt represents a track point.
type Trkpt struct {
	Latitude  float64    `xml:"lat,attr"`
	Longitude float64    `xml:"lon,attr"`
	Elevation *float64   `xml:"ele,omitempty"`
	Time      *time.Time `xml:"time,omitempty"`
}

// Point is a simple latitude/longitude pair.
type Point struct {
	Latitude  float64
	Longitude float64
}

type TimePoint struct {
	Latitude  float64
	Longitude float64
	Time      time.Time
}

func ParseGPX(path string) (*GPX, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	decoder := xml.NewDecoder(f)
	var gpx GPX
	if err := decoder.Decode(&gpx); err != nil && err != io.EOF {
		return nil, err
	}
	return &gpx, nil
}

func GetTrackPoints(gpx *GPX) []Point {
	var pts []Point
	for _, trk := range gpx.Tracks {
		for _, seg := range trk.Segments {
			for _, p := range seg.Points {
				pts = append(pts, Point{Latitude: p.Latitude, Longitude: p.Longitude})
			}
		}
	}
	return pts
}

func GetTrackPointsWithTime(gpx *GPX) []TimePoint {
	var pts []TimePoint
	for _, trk := range gpx.Tracks {
		for _, seg := range trk.Segments {
			for _, p := range seg.Points {
				if p.Time != nil {
					pts = append(pts, TimePoint{Latitude: p.Latitude, Longitude: p.Longitude, Time: *p.Time})
				}
			}
		}
	}
	return pts
}
