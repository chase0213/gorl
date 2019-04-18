package bandit

import (
	"fmt"
	"math"
)

func NewUCB1() *UCB1 {
	return &UCB1{
		arms: []Arm{},
		n:    0,
	}
}

type UCB1 struct {
	arms []Arm

	n int
}

func (b *UCB1) Arms() []Arm {
	return b.arms
}

func (b *UCB1) AddArm(arm Arm) {
	b.arms = append(b.arms, arm)
}

func (b *UCB1) Fit(n int) error {
	delta := 0.01

	for i := 0; i < n; i++ {
		U := make([]float64, len(b.arms))

		for k := 0; k < len(b.arms); k++ {
			// NOTE:
			// Add small number to Ni to avoid divided by 0
			Nk := float64(b.arms[k].N()) + delta

			// Upper confidential band
			U[k] = math.Sqrt(2 * math.Log(float64(b.n)) / Nk)
		}

		j := ArgMaxMeanWithPadding(b.arms, U)
		x := b.arms[j].Pull()
		b.arms[j].Update(x)

		b.n++
	}

	return nil
}

func (b *UCB1) Best() (Arm, error) {
	if b.arms == nil || len(b.arms) == 0 {
		return nil, fmt.Errorf("no arms found")
	}

	j := ArgMaxMean(b.arms)
	return b.arms[j], nil
}

func (b *UCB1) Predict() (float64, error) {
	best, err := b.Best()
	if err != nil {
		return 0.0, err
	}

	return best.Pull(), nil
}

type UCB1Arm struct {
	mean float64
	n    int

	pullFunc PullFunc
}

func (b *UCB1) NewArm(pullFunc PullFunc) *UCB1Arm {
	return &UCB1Arm{
		n: 1,

		pullFunc: pullFunc,
	}
}

func (a *UCB1Arm) N() int {
	return a.n
}

func (a *UCB1Arm) Mean() float64 {
	return a.mean
}

func (a *UCB1Arm) Pull() float64 {
	return a.pullFunc()
}

func (a *UCB1Arm) Update(x float64) {
	a.n++
	n := float64(a.n)
	a.mean = (1-1.0/n)*a.mean + 1.0/n*x
}
