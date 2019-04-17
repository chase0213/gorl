package bandit

import "math/rand"

func generateNormalPullFunc(m float64) func() float64 {
	return func() float64 {
		return rand.NormFloat64() + m
	}
}

// RunExperimentEpsilonGreedy solves multi-armed bandit problem using epsilon greedy algorithm
// ms is initial means for each bandit,
// eps is a threshold to use random or current bandit mean,
// N is the number of trial
func RunExperimentEpsilonGreedy(ms []float64, eps float64, N int) Bandit {
	epsilonGreedy := NewEpsilonGreedy(eps)
	for _, m := range ms {
		pullFunc := generateNormalPullFunc(m)
		arm := epsilonGreedy.NewArm(pullFunc)
		epsilonGreedy.AddArm(arm)
	}
	epsilonGreedy.Fit(N)

	return epsilonGreedy
}

// RunExperimentOptimisticInitialValue solves multi-armed bandit problem using optimistic initial value algorithm
// ms is initial means for each bandit,
// eps is a threshold to use random or current bandit mean,
// N is the number of trial
func RunExperimentOptimisticInitialValue(ms []float64, upperLimit float64, N int) Bandit {
	optimisticInitialValue := NewOptimisticInitialValue(upperLimit)
	for _, m := range ms {
		pullFunc := generateNormalPullFunc(m)
		arm := optimisticInitialValue.NewArm(pullFunc)
		optimisticInitialValue.AddArm(arm)
	}
	optimisticInitialValue.Fit(N)

	return optimisticInitialValue
}
