package fuzzy

import "math"

type AssociativeMemory struct {
	xs, ys Set
	core   [][]*Trapezoid
}

func (am *AssociativeMemory) Init(xs, ys Set) {
	am.xs = xs
	am.ys = ys
	am.core = make([][]*Trapezoid, len(xs))
	for i := range xs {
		am.core[i] = make([]*Trapezoid, len(ys))
	}
}

func (am *AssociativeMemory) Set(t ...*Trapezoid) {
	if len(t) != len(am.core)*len(am.core[0]) {
		panic("Wrong number of trapezoids")
	}

	n := len(am.core)
	for i, tr := range t {
		am.core[i%n][i/n] = tr
	}
}

func (am *AssociativeMemory) Defuzzify(x, y, d float64) float64 {
	xv := am.xs.Value(x)
	yv := am.ys.Value(y)

	val := map[*Trapezoid]float64{}

	for i, vx := range xv {
		for j, vy := range yv {
			if v := math.Min(vx, vy); v > 0 {
				if vv, ok := val[am.core[i][j]]; !ok || v > vv {
					val[am.core[i][j]] = v
				}
			}
		}
	}

	var a, b float64
	s, e := math.MaxFloat64, -math.MaxFloat64
	for tr := range val {
		if tr.LL.X < s {
			s = tr.LL.X
		}
		if tr.RR.X > e {
			e = tr.RR.X
		}
	}
	for x := s; x <= e; x += d {
		f := 0.
		for tr, val := range val {
			f = math.Max(f, math.Min(tr.Val(x), val))
		}
		a += x * f
		b += f
	}

	if b == 0 {
		return (a + b) / 2
	}
	return a / b
}
