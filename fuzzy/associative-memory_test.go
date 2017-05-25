package fuzzy

import (
	"testing"

	"log"

	"github.com/golang/geo/r2"
)

func TestAssociativeMemory_DeFuzzify(t *testing.T) {
	deviationSet := Set{
		{r2.Point{X: -180, Y: 1}, r2.Point{X: -180, Y: 1}, r2.Point{X: -60, Y: 1}, r2.Point{X: -50, Y: 0}},
		{r2.Point{X: -60, Y: 0}, r2.Point{X: -50, Y: 1}, r2.Point{X: -20, Y: 1}, r2.Point{X: -15, Y: 0}},
		{r2.Point{X: -20, Y: 0}, r2.Point{X: -15, Y: 1}, r2.Point{X: 15, Y: 1}, r2.Point{X: 20, Y: 0}},
		{r2.Point{X: 15, Y: 0}, r2.Point{X: 20, Y: 1}, r2.Point{X: 50, Y: 1}, r2.Point{X: 60, Y: 0}},
		{r2.Point{X: 50, Y: 0}, r2.Point{X: 60, Y: 1}, r2.Point{X: 180, Y: 1}, r2.Point{X: 180, Y: 1}},
	}
	tiltSet := Set{
		{r2.Point{X: -90, Y: 1}, r2.Point{X: -90, Y: 1}, r2.Point{X: -30, Y: 1}, r2.Point{X: -25, Y: 0}},
		{r2.Point{X: -30, Y: 0}, r2.Point{X: -25, Y: 1}, r2.Point{X: -10, Y: 1}, r2.Point{X: -5, Y: 0}},
		{r2.Point{X: -10, Y: 0}, r2.Point{X: -5, Y: 1}, r2.Point{X: 5, Y: 1}, r2.Point{X: 10, Y: 0}},
		{r2.Point{X: 5, Y: 0}, r2.Point{X: 10, Y: 1}, r2.Point{X: 25, Y: 1}, r2.Point{X: 30, Y: 0}},
		{r2.Point{X: 25, Y: 0}, r2.Point{X: 30, Y: 1}, r2.Point{X: 90, Y: 1}, r2.Point{X: 90, Y: 1}},
	}

	turnSet := Set{
		{r2.Point{X: -90, Y: 1}, r2.Point{X: -90, Y: 1}, r2.Point{X: -60, Y: 1}, r2.Point{X: -50, Y: 0}},
		{r2.Point{X: -60, Y: 0}, r2.Point{X: -50, Y: 1}, r2.Point{X: -20, Y: 1}, r2.Point{X: -15, Y: 0}},
		{r2.Point{X: -20, Y: 0}, r2.Point{X: -15, Y: 1}, r2.Point{X: 15, Y: 1}, r2.Point{X: 20, Y: 0}},
		{r2.Point{X: 15, Y: 0}, r2.Point{X: 20, Y: 1}, r2.Point{X: 50, Y: 1}, r2.Point{X: 60, Y: 0}},
		{r2.Point{X: 50, Y: 0}, r2.Point{X: 60, Y: 1}, r2.Point{X: 90, Y: 1}, r2.Point{X: 90, Y: 1}},
	}

	am := AssociativeMemory{}
	am.Init(deviationSet, tiltSet)
	am.Set(
		&turnSet[0], &turnSet[0], &turnSet[1], &turnSet[2], &turnSet[2],
		&turnSet[0], &turnSet[1], &turnSet[1], &turnSet[2], &turnSet[3],
		&turnSet[1], &turnSet[1], &turnSet[2], &turnSet[3], &turnSet[3],
		&turnSet[1], &turnSet[1], &turnSet[3], &turnSet[3], &turnSet[4],
		&turnSet[1], &turnSet[2], &turnSet[3], &turnSet[4], &turnSet[4],
	)

	log.Println(am.Defuzzify(-20, 0, 0.5))
}
