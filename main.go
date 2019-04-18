package main

import (
	crand "crypto/rand"
	"time"

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

	Ns := []int{100, 10000, 1000000}

	var beginAt time.Time
	var endAt time.Time

	for _, N := range Ns {
		fmt.Printf("\n=== Epsilon Greedy Algorithm ===\n")
		beginAt = time.Now()
		banditEpsilonGreedy := bandit.RunExperimentEpsilonGreedy(ms, eps, N)
		endAt = time.Now()
		fmt.Printf("Time %f [sec] for %d-iterations\n", endAt.Sub(beginAt).Seconds(), N)
		for i, a := range banditEpsilonGreedy.Arms() {
			fmt.Printf("True mean of bandit(No. %d) = %f, actual = %f\n", i+1, ms[i], a.Mean())
		}

		fmt.Printf("\n=== Optimistic Initial Value Algorithm ===\n")
		beginAt = time.Now()
		banditOptimisticInitialValue := bandit.RunExperimentOptimisticInitialValue(ms, upperLimit, N)
		endAt = time.Now()
		fmt.Printf("Time %f [sec] for %d-iterations\n", endAt.Sub(beginAt).Seconds(), N)
		for i, a := range banditOptimisticInitialValue.Arms() {
			fmt.Printf("True mean of bandit(No. %d) = %f, actual = %f\n", i+1, ms[i], a.Mean())
		}

		fmt.Printf("\n=== UCB1 Algorithm ===\n")
		beginAt = time.Now()
		ucb1 := bandit.RunExperimentUCB1(ms, N)
		endAt = time.Now()
		fmt.Printf("Time %f [sec] for %d-iterations\n", endAt.Sub(beginAt).Seconds(), N)
		for i, a := range ucb1.Arms() {
			fmt.Printf("True mean of bandit(No. %d) = %f, actual = %f\n", i+1, ms[i], a.Mean())
		}
	}

}
