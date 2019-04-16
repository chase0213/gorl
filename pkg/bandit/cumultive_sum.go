package bandit

func CumultiveSum(data []float64) []float64 {
	s := make([]float64, len(data))
	for i, x := range data {
		if i == 0 {
			s[i] = x
		} else {
			s[i] = s[i-1] + x
		}
	}
	return s
}
