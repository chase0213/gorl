package bandit

import "fmt"

func NewOptimisticInitialValue(upperLimit float64) *OptimisticInitialValue {
	return &OptimisticInitialValue{
		arms:       []Arm{},
		UpperLimit: upperLimit,
	}
}

type OptimisticInitialValue struct {
	arms []Arm

	UpperLimit float64 `json:"upper_limit"`
}

func (b *OptimisticInitialValue) Arms() []Arm {
	return b.arms
}

func (b *OptimisticInitialValue) AddArm(arm Arm) {
	b.arms = append(b.arms, arm)
}

func (b *OptimisticInitialValue) Fit(n int) error {
	for i := 0; i < n; i++ {
		j := ArgMaxMean(b.arms)
		x := b.arms[j].Pull()
		b.arms[j].Update(x)
	}

	return nil
}

func (b *OptimisticInitialValue) Best() (Arm, error) {
	if b.arms == nil || len(b.arms) == 0 {
		return nil, fmt.Errorf("no arms found")
	}

	j := ArgMaxMean(b.arms)
	return b.arms[j], nil
}

func (b *OptimisticInitialValue) Predict() (float64, error) {
	best, err := b.Best()
	if err != nil {
		return 0.0, err
	}

	return best.Pull(), nil
}

type OptimisticInitialValueArm struct {
	mean float64
	n    int

	pullFunc PullFunc
}

func (b *OptimisticInitialValue) NewArm(pullFunc PullFunc) *OptimisticInitialValueArm {
	return &OptimisticInitialValueArm{
		mean: b.UpperLimit,
		n:    1,

		pullFunc: pullFunc,
	}
}

func (a *OptimisticInitialValueArm) N() int {
	return a.n
}

func (a *OptimisticInitialValueArm) Mean() float64 {
	return a.mean
}

func (a *OptimisticInitialValueArm) Pull() float64 {
	return a.pullFunc()
}

func (a *OptimisticInitialValueArm) Update(x float64) {
	a.n++
	n := float64(a.n)
	a.mean = (1-1.0/n)*a.mean + 1.0/n*x
}
