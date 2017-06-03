package main

import (
	"math"

	"github.com/golang/geo/r3"
)

type Surface struct {
	core func(float64, float64) float64
}

func (s *Surface) SetCore(core func(float64, float64) float64) {
	s.core = core
}

func (s *Surface) GetHeight(x, y float64) float64 {
	return s.core(x, y)
}

// This function calculates angle proection of normalized vector at left side and right side
func (s *Surface) GetVehicleAngels(x, y, phi, a, b float64) (float64, float64) {
	xbr := x + a*math.Cos(phi) + b*math.Sin(phi)
	xtr := x + a*math.Cos(phi) - b*math.Sin(phi)
	xtl := x - a*math.Cos(phi) - b*math.Sin(phi)
	xbl := x - a*math.Cos(phi) + b*math.Sin(phi)

	ybr := y + a*math.Sin(phi) - b*math.Cos(phi)
	ytr := y + a*math.Sin(phi) + b*math.Cos(phi)
	ytl := y - a*math.Sin(phi) + b*math.Cos(phi)
	ybl := y - a*math.Sin(phi) - b*math.Cos(phi)

	ztr := s.GetHeight(xtr, ytr)
	ztl := s.GetHeight(xtl, ytl)
	zbl := s.GetHeight(xbl, ybl)
	zbr := s.GetHeight(xbr, ybr)

	tltr := r3.Vector{X: xtr - xtl, Y: ytr - ytl, Z: ztr - ztl}
	trbr := r3.Vector{X: xbr - xtr, Y: ybr - ytr, Z: zbr - ztr}
	brbl := r3.Vector{X: xbl - xbr, Y: ybl - ybr, Z: zbl - zbr}
	bltl := r3.Vector{X: xtl - xbl, Y: ytl - ybl, Z: ztl - zbl}

	ntl := bltl.Cross(tltr).Normalize()
	ntr := tltr.Cross(trbr).Normalize()
	nbr := trbr.Cross(brbl).Normalize()
	nbl := brbl.Cross(bltl).Normalize()

	n := ntl.Add(ntr).Add(nbr).Add(nbl).Mul(0.25)

	return 90 - n.Angle(r3.Vector{X: -math.Sin(phi), Y: math.Cos(phi), Z: 0}).Degrees(), -90 + n.Angle(r3.Vector{X: math.Cos(phi), Y: math.Sin(phi), Z: 0}).Degrees()
}
