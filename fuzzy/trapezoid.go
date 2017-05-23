package fuzzy

import "github.com/golang/geo/r2"

type Trapezoid struct {
	LL, L, R, RR r2.Point
}

func (tr *Trapezoid) Init(LL, L, R, RR r2.Point) {
	if LL.X > L.X || L.X > R.X || R.X > RR.X {
		panic("Wrong sequince of points")
	}

	tr.LL = LL
	tr.L = L
	tr.R = R
	tr.RR = RR
}

func (tr *Trapezoid) Val(x float64) float64{
	var y float64
	if tr.LL.X <= x && x < tr.L.X {
		y = tr.LL.Y + (x-tr.LL.X)/(tr.L.X-tr.LL.X)*(tr.L.Y-tr.LL.Y)
	}
	if tr.L.X <= x && x < tr.R.X {
		y = tr.L.Y + (x-tr.L.X)/(tr.R.X-tr.L.X)*(tr.R.Y-tr.L.Y)
	}
	if tr.R.X <= x && x <= tr.RR.X {
		y = tr.R.Y + (x-tr.R.X)/(tr.RR.X-tr.R.X)*(tr.RR.Y-tr.R.Y)
	}
	return y
}