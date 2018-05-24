package main

import (
	"math"
)

// ScoreTraectory -> min
func ScoreTraectory(vehicle Vehicle, traectory []VehiclePosition, surface Surface) float64 {
	var score float64

	for _, pos := range traectory {
		vehicle.ApplyPosition(&pos)

		_, tilt := vehicle.FindDegrees(surface.GetHeight)

		score += math.Abs(tilt)
		//fmt.Println(tang, tilt)
	}

	return score
}
