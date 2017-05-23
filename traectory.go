package main

import (
	"math"

	"log"

	"github.com/ZzEeKkAa/fuzzy-logic/fuzzy"
	"github.com/golang/geo/r2"
	"github.com/golang/geo/r3"
)

func BuildTraectory(a, b float64, start r2.Point, startAngel float64, end r2.Point, d float64, am fuzzy.AssociativeMemory, s Surface) []struct {
	r2.Point
	float64
} {
	var ans []struct {
		r2.Point
		float64
	}

	ans = append(ans, struct {
		r2.Point
		float64
	}{start, startAngel})
	angel := r3.Vector{X: math.Cos(startAngel / 90 * math.Pi), Y: math.Sin(startAngel / 90 * math.Pi), Z: 0}
	destination := r3.Vector{X: end.X, Y: end.Y, Z: 0}
	floatAngel := startAngel
	max := 100
	for p := start; end.Sub(p).Norm() > d && max > 0; max-- {
		log.Println(p)
		_, tilt := s.GetVehicleAngels(p.X, p.Y, startAngel, a, b)
		deviation := destination.Angle(angel).Degrees()
		turn := am.Defuzzify(deviation, tilt, d)
		floatAngel += turn
		angel = r3.Vector{X: math.Cos(floatAngel / 90 * math.Pi), Y: math.Sin(floatAngel / 90 * math.Pi), Z: 0}
		floatAngel = normalizeAngel(floatAngel)
		mv := angel.Normalize().Mul(d)
		p.X += mv.X
		p.Y += mv.Y
	}

	return ans
}

func normalizeAngel(a float64) float64 {
	a -= float64(int(a/360)) * 360
	if a > 180 {
		a -= 360
	}
	if a < -180 {
		a += 360
	}
	return a
}
