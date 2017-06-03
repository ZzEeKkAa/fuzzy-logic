package main

import (
	"math"

	"log"

	"github.com/golang/geo/r3"
)

type Vehicle struct {
	x, y  float64
	a, b  float64
	alpha float64
}

// tang, tilt
func (v *Vehicle) FindDegrees(h func(float64, float64) float64) (float64, float64) {
	alpha := v.alpha / 180 * math.Pi
	xbr := v.x + v.a*math.Cos(alpha) + v.b*math.Sin(alpha)
	xtr := v.x + v.a*math.Cos(alpha) - v.b*math.Sin(alpha)
	xtl := v.x - v.a*math.Cos(alpha) - v.b*math.Sin(alpha)
	xbl := v.x - v.a*math.Cos(alpha) + v.b*math.Sin(alpha)

	ybr := v.y + v.a*math.Sin(alpha) - v.b*math.Cos(alpha)
	ytr := v.y + v.a*math.Sin(alpha) + v.b*math.Cos(alpha)
	ytl := v.y - v.a*math.Sin(alpha) + v.b*math.Cos(alpha)
	ybl := v.y - v.a*math.Sin(alpha) - v.b*math.Cos(alpha)

	ztr := h(xtr, ytr)
	ztl := h(xtl, ytl)
	zbl := h(xbl, ybl)
	zbr := h(xbr, ybr)

	tltr := r3.Vector{X: xtr - xtl, Y: ytr - ytl, Z: ztr - ztl}
	trbr := r3.Vector{X: xbr - xtr, Y: ybr - ytr, Z: zbr - ztr}
	brbl := r3.Vector{X: xbl - xbr, Y: ybl - ybr, Z: zbl - zbr}
	bltl := r3.Vector{X: xtl - xbl, Y: ytl - ybl, Z: ztl - zbl}

	ntl := bltl.Cross(tltr).Normalize()
	ntr := tltr.Cross(trbr).Normalize()
	nbr := trbr.Cross(brbl).Normalize()
	nbl := brbl.Cross(bltl).Normalize()

	n := ntl.Add(ntr).Add(nbr).Add(nbl).Mul(0.25)

	return 90 - n.Angle(r3.Vector{X: -math.Sin(alpha), Y: math.Cos(alpha), Z: 0}).Degrees(), -90 + n.Angle(r3.Vector{X: math.Cos(alpha), Y: math.Sin(alpha), Z: 0}).Degrees()
}

func (v *Vehicle) TurnRight(dAlpha float64) {
	v.alpha -= dAlpha
}

func (v *Vehicle) TurnLeft(dAlpha float64) {
	v.TurnRight(-dAlpha)
}

func (v *Vehicle) MoveForward(dDist float64) {
	alpha := v.alpha / 180 * math.Pi
	v.x += -dDist * math.Sin(alpha)
	v.y += dDist * math.Cos(alpha)
}

func (v *Vehicle) FindDeviation(x, y float64) float64 {
	x -= v.x
	y -= v.y
	log.Println("Dev:", x, y)
	ang := r3.Vector{x, y, 0}.Angle(r3.Vector{0, 1, 0}).Degrees()
	log.Println("Dev_ang", ang)
	if x > 0 {
		ang = -ang
	}
	log.Println("Dev_ang", ang)

	ang -= v.alpha //- ang

	for ang > 180 {
		ang -= 360
	}
	for ang < -180 {
		ang += 360
	}

	return -ang
}

type VehiclePosition struct {
	x, y, alpha float64
}

func Pos(v *Vehicle) VehiclePosition {
	return VehiclePosition{x: v.x, y: v.y, alpha: v.alpha}
}
