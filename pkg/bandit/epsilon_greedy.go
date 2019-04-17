package bandit

type EpsilonGreedy struct {
	mean float64
	n    int

	pullFunc PullFunc
}

func NewEpsilonGreedy(pullFunc PullFunc) *EpsilonGreedy {
	return &EpsilonGreedy{
		mean:     0.0,
		n:        0,
		pullFunc: pullFunc,
	}
}

func (b *EpsilonGreedy) N() int {
	return b.n
}

func (b *EpsilonGreedy) Mean() float64 {
	return b.mean
}

func (b *EpsilonGreedy) Pull() float64 {
	return b.pullFunc()
}

func (b *EpsilonGreedy) Update(x float64) {
	b.n++
	n := float64(b.n)
	b.mean = (1-1.0/n)*b.mean + 1.0/n*x
}
