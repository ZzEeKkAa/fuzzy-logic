package main

import (
	"log"
	"math"
	"testing"
)

func TestSurface_GetVehicleAngels(t *testing.T) {
	x, y := 0., 0.
	a, b := 1., 2.
	phi := 90. / 180. * math.Pi

	xbr := x + a*math.Cos(phi) + b*math.Sin(phi)
	xtr := x + a*math.Cos(phi) - b*math.Sin(phi)
	xtl := x - a*math.Cos(phi) - b*math.Sin(phi)
	xbl := x - a*math.Cos(phi) + b*math.Sin(phi)

	ybr := y + a*math.Sin(phi) - b*math.Cos(phi)
	ytr := y + a*math.Sin(phi) + b*math.Cos(phi)
	ytl := y - a*math.Sin(phi) + b*math.Cos(phi)
	ybl := y - a*math.Sin(phi) - b*math.Cos(phi)

	log.Println(xtl, ytl)
	log.Println(xtr, ytr)
	log.Println(xbl, ybl)
	log.Println(xbr, ybr)
}
