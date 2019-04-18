package main

import (
	crand "crypto/rand"

	"fmt"
	"math"
	"math/big"
	"math/rand"

	"github.com/chase0213/gorl/pkg/bandit"
)

func main() {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())

	ms := []float64{
		0.1, 0.2, 0.3,
	}
	eps := 0.1
	upperLimit := 10.0
	N := 10000

	fmt.Printf("\n=== Epsilon Greedy Algorithm ===\n")
	banditEpsilonGreedy := bandit.RunExperimentEpsilonGreedy(ms, eps, N)
	for i, a := range banditEpsilonGreedy.Arms() {
		fmt.Printf("True mean of bandit(No. %d) = %f, actual = %f\n", i+1, ms[i], a.Mean())
	}

	fmt.Printf("\n=== Optimistic Initial Value Algorithm ===\n")
	banditOptimisticInitialValue := bandit.RunExperimentOptimisticInitialValue(ms, upperLimit, N)
	for i, a := range banditOptimisticInitialValue.Arms() {
		fmt.Printf("True mean of bandit(No. %d) = %f, actual = %f\n", i+1, ms[i], a.Mean())
	}

}
