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

	// Vehicle angel in two systems
	floatAngel := startAngel
	angel := r3.Vector{X: math.Cos(startAngel / 180 * math.Pi), Y: math.Sin(startAngel / 180 * math.Pi), Z: 0}

	destination := r3.Vector{X: end.X, Y: end.Y, Z: 0}
	max := 10000
	for p := start; end.Sub(p).Norm() > d && max > 0; max-- {
		log.Println(p)
		log.Println("Current Pos:", p)
		log.Println("Current Ang:", floatAngel, r3.Vector{X: 1, Y: 0, Z: 0}.Angle(angel).Degrees())
		log.Println("Dist left:", end.Sub(p).Norm())

		_, tilt := s.GetVehicleAngels(p.X, p.Y, floatAngel, a, b)
		deviation := destination.Sub(r3.Vector{X: p.X, Y: p.Y, Z: 0}).Angle(angel).Degrees()
		deviation = normalizeAngel(deviation)

		log.Println("til:", tilt)
		log.Println("dev:", deviation)

		turn := am.Defuzzify(deviation, tilt, d)
		log.Println("Turn:", turn)

		floatAngel += turn
		angel = r3.Vector{X: math.Cos(floatAngel / 180 * math.Pi), Y: math.Sin(floatAngel / 180 * math.Pi), Z: 0}
		floatAngel = normalizeAngel(floatAngel)
		mv := angel.Normalize().Mul(d)
		log.Println(angel.Normalize())
		p.X += mv.X
		p.Y += mv.Y
		log.Println()

		ans = append(ans, struct {
			r2.Point
			float64
		}{p, floatAngel})
	}
	if max == 0 {
		log.Println("Steps limit reached")
	} else {
		log.Println("End point reached. Number of steps", 1000-max)
	}

	return ans
}

func normalizeAngel(a float64) float64 {
	//a -= float64(int(a/360)) * 360
	//if a > 360 {
	//	a -= 360
	//}
	//if a < 0 {
	//	a += 360
	//}
	if a > 180 {
		a -= 360
	}
	if a < -180 {
		a += 360
	}
	return a
}

func BuildTrajectoryNew(v Vehicle, end r2.Point, moveStep float64, eps float64, amUp fuzzy.AssociativeMemory, amDown fuzzy.AssociativeMemory, h func(float64, float64) float64) (ans []VehiclePosition) {
	d := 0.01
	ans = append(ans, VehiclePosition{v.x, v.y, v.alpha})

	var defuzzify = amDown.Defuzzify
	var alpha = 15.

	max := 100000
	for ; end.Sub(r2.Point{v.x, v.y}).Norm() > eps && max > 0; max-- {
		tang, tilt := v.FindDegrees(h)
		deviation := v.FindDeviation(end.X, end.Y)

		log.Println(v.x, v.y)

		log.Println("til:", tilt)
		log.Println("dev:", deviation)

		var turn float64
		if tang < -alpha {
			defuzzify = amDown.Defuzzify
		} else if tang > alpha {
			defuzzify = amUp.Defuzzify
		}
		turn = defuzzify(deviation, tilt, d)

		log.Println("Turn:", turn)

		log.Println("Angel", v.alpha)

		v.TurnRight(turn)
		v.MoveForward(moveStep)

		ans = append(ans, Pos(&v))
		log.Println()
	}
	if max == 0 {
		log.Println("Steps limit reached")
	} else {
		log.Println("End point reached. Number of steps", 10000-max)
	}

	return
}
