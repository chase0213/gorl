package bandit

import "math/rand"

func generateNormalPullFunc(m float64) func() float64 {
	return func() float64 {
		return rand.NormFloat64() + m
	}
}

// RunEpsilonGreedy solves multi-armed bandit problem using epsilon greedy algorithm
// ms is initial means for each bandit,
// eps is a threshold to use random or current bandit mean,
// N is the number of trial
func RunEpsilonGreedy(ms []float64, eps float64, N int) ([]Bandit, []float64) {
	bandits := make([]Bandit, len(ms))
	for i, m := range ms {
		pullFunc := generateNormalPullFunc(m)
		bandits[i] = NewEpsilonGreedy(pullFunc)
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

// RunOptimisticInitialValue solves multi-armed bandit problem using optimistic initial value algorithm
// ms is initial means for each bandit,
// eps is a threshold to use random or current bandit mean,
// N is the number of trial
func RunOptimisticInitialValue(ms []float64, upperLimit float64, N int) ([]Bandit, []float64) {
	bandits := make([]Bandit, len(ms))
	for i, m := range ms {
		pullFunc := generateNormalPullFunc(m)
		bandits[i] = NewOptimisticInitialValue(pullFunc, upperLimit)
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
