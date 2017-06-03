package main

import (
	"math"
	"testing"
)

func TestVehicle_FindDegrees(t *testing.T) {
	v := Vehicle{x: 0, y: 0, a: 1, b: 2, alpha: 0}

	ang := 180.0
	ang *= math.Pi / 180

	t.Log(v.FindDeviation(math.Cos(ang), math.Sin(ang)))

	v = Vehicle{x: 0, y: 0, a: 1, b: 2, alpha: 90}

	t.Log(Pos(&v))

	v.MoveForward(5)
	v.TurnLeft(45)
	v.MoveForward(2)

	t.Log(Pos(&v))

	tang, tilt := v.FindDegrees(func(x float64, y float64) float64 {
		return 2 * (1 + y/3)
	})

	t.Log(tang, tilt)
}
