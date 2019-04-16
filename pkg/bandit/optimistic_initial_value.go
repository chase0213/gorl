package bandit

import (
	"math/rand"
)

type OptimisticInitialValue struct {
	m    float64
	mean float64
	n    int
}

func NewOptimisticInitialValue(m float64, upperLimit float64) *OptimisticInitialValue {
	return &OptimisticInitialValue{
		m:    m,
		mean: upperLimit,
		n:    1,
	}
}

func (b *OptimisticInitialValue) N() int {
	return b.n
}

func (b *OptimisticInitialValue) M() float64 {
	return b.m
}

func (b *OptimisticInitialValue) Mean() float64 {
	return b.mean
}

func (b *OptimisticInitialValue) Pull() float64 {
	return rand.NormFloat64() + b.m
}

func (b *OptimisticInitialValue) Update(x float64) {
	b.n++
	n := float64(b.n)
	b.mean = (1-1.0/n)*b.mean + 1.0/n*x
}

// RunOptimisticInitialValue solves multi-armed bandit problem using optimistic initial value algorithm
// ms is initial means for each bandit,
// eps is a threshold to use random or current bandit mean,
// N is the number of trial
func RunOptimisticInitialValue(ms []float64, upperLimit float64, N int) ([]Bandit, []float64) {
	bandits := make([]Bandit, len(ms))
	for i := range ms {
		bandits[i] = NewOptimisticInitialValue(ms[i], upperLimit)
	}

	data := make([]float64, N)

	for i := 0; i < N; i++ {
		j := ArgMaxMean(bandits)
		x := bandits[j].Pull()
		bandits[j].Update(x)

		// データプロット用
		data[i] = x
	}

	return bandits, data
}
