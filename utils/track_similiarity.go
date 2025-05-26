package utils

import (
	"math"
	"time"
)

const earthRadius = 6371000.0 // meters

func distancePointToSegment(tp TimePoint, A, B Point) float64 {
	latAvg := (A.Latitude + B.Latitude) / 2 * math.Pi / 180

	dLon := (B.Longitude - A.Longitude) * math.Pi / 180

	dLat := (B.Latitude - A.Latitude) * math.Pi / 180

	ABx := earthRadius * dLon * math.Cos(latAvg)
	ABy := earthRadius * dLat

	dLonP := (tp.Longitude - A.Longitude) * math.Pi / 180
	dLatP := (tp.Latitude - A.Latitude) * math.Pi / 180

	APx := earthRadius * dLonP * math.Cos(latAvg)
	APy := earthRadius * dLatP

	ab2 := ABx*ABx + ABy*ABy
	if ab2 == 0 {
		return math.Hypot(APx, APy)
	}

	t := (APx*ABx + APy*ABy) / ab2
	var cx, cy float64
	switch {
	case t <= 0:
		cx, cy = 0, 0
	case t >= 1:
		cx, cy = ABx, ABy
	default:
		cx, cy = ABx*t, ABy*t
	}

	dx := APx - cx
	dy := APy - cy
	return math.Hypot(dx, dy)
}

func ContainsRouteSlidingWindow(timePoints []TimePoint, route []Point, tol float64) (bool, time.Duration) {
	n := len(route)
	if n < 2 {
		return false, 0
	}
	segCount := n - 1

	matched := make([]bool, segCount)
	matchTimes := make([]time.Time, segCount)

	tpIndex := 0
	for i := 0; i < segCount && tpIndex < len(timePoints); i++ {
		A, B := route[i], route[i+1]
		for j := tpIndex; j < len(timePoints); j++ {
			d := distancePointToSegment(timePoints[j], A, B)
			if d <= tol {
				matched[i] = true
				matchTimes[i] = timePoints[j].Time
				tpIndex = j
				break
			}
		}
	}

	window := 5
	for start := 0; start <= segCount-window; start++ {
		missCount := 0
		for k := start; k < start+window; k++ {
			if !matched[k] {
				missCount++
				if missCount > 1 {
					return false, 0
				}
			}
		}
	}

	var firstTime, lastTime time.Time
	for i, ok := range matched {
		if ok {
			firstTime = matchTimes[i]
			break
		}
	}
	for i := segCount - 1; i >= 0; i-- {
		if matched[i] {
			lastTime = matchTimes[i]
			break
		}
	}

	if firstTime.IsZero() || lastTime.IsZero() {
		return false, 0
	}
	return true, lastTime.Sub(firstTime)
}
