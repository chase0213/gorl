package bandit

import (
	"math/rand"
)

type EpsilonGreedy struct {
	m    float64
	mean float64
	n    int
}

func NewEpsilonGreedy(m float64) *EpsilonGreedy {
	return &EpsilonGreedy{
		m:    m,
		mean: 0.0,
		n:    0,
	}
}

func (b *EpsilonGreedy) N() int {
	return b.n
}

func (b *EpsilonGreedy) M() float64 {
	return b.m
}

func (b *EpsilonGreedy) Mean() float64 {
	return b.mean
}

func (b *EpsilonGreedy) Pull() float64 {
	return rand.NormFloat64() + b.m
}

func (b *EpsilonGreedy) Update(x float64) {
	b.n++
	n := float64(b.n)
	b.mean = (1-1.0/n)*b.mean + 1.0/n*x
}

// RunEpsilonGreedy solves multi-armed bandit problem using epsilon greedy algorithm
// ms is initial means for each bandit,
// eps is a threshold to use random or current bandit mean,
// N is the number of trial
func RunEpsilonGreedy(ms []float64, eps float64, N int) ([]Bandit, []float64) {
	bandits := make([]Bandit, len(ms))
	for i := range ms {
		bandits[i] = NewEpsilonGreedy(ms[i])
	}

	data := make([]float64, N)

	var j int
	for i := 0; i < N; i++ {
		p := rand.NormFloat64()
		if p < eps {
			// サイコロの目が eps より小さければ、無作為に選ぶ
			j = rand.Int() % len(ms)
		} else {
			// サイコロの目が eps より大きければ、Bandit の中から一番大きいものを選ぶ
			j = ArgMaxMean(bandits)
		}

		x := bandits[j].Pull()
		bandits[j].Update(x)

		// データプロット用
		data[i] = x
	}

	return bandits, data
}
