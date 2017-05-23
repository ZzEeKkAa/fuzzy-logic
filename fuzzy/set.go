package fuzzy

type Set []Trapezoid

func (s Set) Value(x float64) []float64 {
	res := make([]float64, 0, len(s))
	for _, tr := range s {
		res = append(res, tr.Val(x))
	}
	return res
}
