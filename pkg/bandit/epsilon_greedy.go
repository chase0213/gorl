package bandit

import (
	"fmt"
	"math/rand"
)

func NewEpsilonGreedy(eps float64) *EpsilonGreedy {
	return &EpsilonGreedy{
		arms: []Arm{},
		Eps:  eps,
	}
}

type EpsilonGreedy struct {
	arms []Arm

	Eps float64 `json:"eps"`
}

func (b *EpsilonGreedy) Arms() []Arm {
	return b.arms
}

func (b *EpsilonGreedy) AddArm(arm Arm) {
	b.arms = append(b.arms, arm)
}

func (b *EpsilonGreedy) Fit(n int) error {
	// the number of arms
	m := len(b.arms)

	var j int
	for i := 0; i < n; i++ {
		p := rand.NormFloat64()
		if p < b.Eps {
			// サイコロの目が eps より小さければ、無作為に選ぶ
			j = rand.Int() % m
		} else {
			// サイコロの目が eps より大きければ、Bandit の中から一番大きいものを選ぶ
			j = ArgMaxMean(b.arms)
		}

		x := b.arms[j].Pull()
		b.arms[j].Update(x)
	}

	return nil
}

func (b *EpsilonGreedy) Best() (Arm, error) {
	if b.arms == nil || len(b.arms) == 0 {
		return nil, fmt.Errorf("no arms found")
	}

	j := ArgMaxMean(b.arms)
	return b.arms[j], nil
}

func (b *EpsilonGreedy) Predict() (float64, error) {
	best, err := b.Best()
	if err != nil {
		return 0.0, err
	}

	return best.Pull(), nil
}

type EpsilonGreedyArm struct {
	mean float64
	n    int

	pullFunc PullFunc
}

func (b *EpsilonGreedy) NewArm(pullFunc PullFunc) *EpsilonGreedyArm {
	return &EpsilonGreedyArm{
		mean:     0.0,
		n:        0,
		pullFunc: pullFunc,
	}
}

func (a *EpsilonGreedyArm) N() int {
	return a.n
}

func (a *EpsilonGreedyArm) Mean() float64 {
	return a.mean
}

func (a *EpsilonGreedyArm) Pull() float64 {
	return a.pullFunc()
}

func (a *EpsilonGreedyArm) Update(x float64) {
	a.n++
	n := float64(a.n)
	a.mean = (1-1.0/n)*a.mean + 1.0/n*x
}
