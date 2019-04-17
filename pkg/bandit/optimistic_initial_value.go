package bandit

type OptimisticInitialValue struct {
	mean float64
	n    int

	pullFunc PullFunc
}

func NewOptimisticInitialValue(pullFunc PullFunc, upperLimit float64) *OptimisticInitialValue {
	return &OptimisticInitialValue{
		mean: upperLimit,
		n:    1,

		pullFunc: pullFunc,
	}
}

func (b *OptimisticInitialValue) N() int {
	return b.n
}

func (b *OptimisticInitialValue) Mean() float64 {
	return b.mean
}

func (b *OptimisticInitialValue) Pull() float64 {
	return b.pullFunc()
}

func (b *OptimisticInitialValue) Update(x float64) {
	b.n++
	n := float64(b.n)
	b.mean = (1-1.0/n)*b.mean + 1.0/n*x
}
