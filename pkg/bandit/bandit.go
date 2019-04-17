package bandit

// Bandit is an interface of solver for bandit problem
type Bandit interface {
	Fit(n int) error
	Predict() (float64, error)

	Arms() []Arm
	AddArm(Arm)
	Best() (Arm, error)
}

// Arm is an interface of multi-armed bandit problem
type Arm interface {
	N() int
	Mean() float64

	Pull() float64
	Update(x float64)
}

type PullFunc func() float64

func ArgMaxMean(arms []Arm) int {
	if len(arms) == 0 {
		return -1
	}

	idx := 0
	value := arms[idx].Mean()
	for i := 1; i < len(arms); i++ {
		if arms[i].Mean() >= value {
			value = arms[i].Mean()
			idx = i
		}
	}
	return idx
}
