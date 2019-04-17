package bandit

// Bandit is an interface of multi-armed bandit problem
type Bandit interface {
	N() int
	Mean() float64

	Pull() float64
	Update(x float64)
}

type PullFunc func() float64

func ArgMaxMean(bandits []Bandit) int {
	if len(bandits) == 0 {
		return -1
	}

	idx := 0
	value := bandits[idx].Mean()
	for i := 1; i < len(bandits); i++ {
		if bandits[i].Mean() >= value {
			value = bandits[i].Mean()
			idx = i
		}
	}
	return idx
}
