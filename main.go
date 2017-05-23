package main

import (
	"log"
	"math"

	"github.com/ZzEeKkAa/fuzzy-logic/fuzzy"
	"github.com/golang/geo/r2"
)

func main() {
	var s Surface
	s.SetCore(func(x, y float64) float64 {
		return math.Sin(x * y)
	})
	//for i := 0.; i < 10; i += 0.01 {
	//	log.Println(s.GetVehicleAngels(i, 2*i, 0, 0.05, 0.02))
	//}

	deviationSet := fuzzy.Set{
		{r2.Point{X: -180, Y: 1}, r2.Point{X: -180, Y: 1}, r2.Point{X: -60, Y: 1}, r2.Point{X: -50, Y: 0}},
		{r2.Point{X: -60, Y: 0}, r2.Point{X: -50, Y: 1}, r2.Point{X: -20, Y: 1}, r2.Point{X: -15, Y: 0}},
		{r2.Point{X: -20, Y: 0}, r2.Point{X: -15, Y: 1}, r2.Point{X: 15, Y: 1}, r2.Point{X: 20, Y: 0}},
		{r2.Point{X: 15, Y: 0}, r2.Point{X: 20, Y: 1}, r2.Point{X: 50, Y: 1}, r2.Point{X: 60, Y: 0}},
		{r2.Point{X: 50, Y: 0}, r2.Point{X: 60, Y: 1}, r2.Point{X: 180, Y: 1}, r2.Point{X: 180, Y: 1}},
	}
	tiltSet := fuzzy.Set{
		{r2.Point{X: -90, Y: 1}, r2.Point{X: -90, Y: 1}, r2.Point{X: -30, Y: 1}, r2.Point{X: -25, Y: 0}},
		{r2.Point{X: -30, Y: 0}, r2.Point{X: -25, Y: 1}, r2.Point{X: -10, Y: 1}, r2.Point{X: -5, Y: 0}},
		{r2.Point{X: -10, Y: 0}, r2.Point{X: -5, Y: 1}, r2.Point{X: 5, Y: 1}, r2.Point{X: 10, Y: 0}},
		{r2.Point{X: 5, Y: 0}, r2.Point{X: 10, Y: 1}, r2.Point{X: 25, Y: 1}, r2.Point{X: 30, Y: 0}},
		{r2.Point{X: 25, Y: 0}, r2.Point{X: 30, Y: 1}, r2.Point{X: 90, Y: 1}, r2.Point{X: 90, Y: 1}},
	}

	turnSet := fuzzy.Set{
		{r2.Point{X: -180, Y: 1}, r2.Point{X: -180, Y: 1}, r2.Point{X: -60, Y: 1}, r2.Point{X: -50, Y: 0}},
		{r2.Point{X: -60, Y: 0}, r2.Point{X: -50, Y: 1}, r2.Point{X: -20, Y: 1}, r2.Point{X: -15, Y: 0}},
		{r2.Point{X: -20, Y: 0}, r2.Point{X: -15, Y: 1}, r2.Point{X: 15, Y: 1}, r2.Point{X: 20, Y: 0}},
		{r2.Point{X: 15, Y: 0}, r2.Point{X: 20, Y: 1}, r2.Point{X: 50, Y: 1}, r2.Point{X: 60, Y: 0}},
		{r2.Point{X: 50, Y: 0}, r2.Point{X: 60, Y: 1}, r2.Point{X: 180, Y: 1}, r2.Point{X: 180, Y: 1}},
	}

	am := fuzzy.AssociativeMemory{}
	am.Init(deviationSet, tiltSet)
	am.Set(
		&turnSet[0], &turnSet[0], &turnSet[1], &turnSet[2], &turnSet[2],
		&turnSet[0], &turnSet[1], &turnSet[1], &turnSet[2], &turnSet[3],
		&turnSet[1], &turnSet[1], &turnSet[2], &turnSet[3], &turnSet[3],
		&turnSet[1], &turnSet[1], &turnSet[3], &turnSet[3], &turnSet[4],
		&turnSet[1], &turnSet[2], &turnSet[3], &turnSet[4], &turnSet[4],
	)

	traectory := BuildTraectory(0.05, 0.02, r2.Point{X: 0, Y: 0}, 0, r2.Point{X: 20, Y: 10}, 0.01, am, s)
	log.Println(traectory)
}
