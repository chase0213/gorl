# gorl

gorl is a go language implementation of reinforcement learning.

This package is under development.
You shouldn't use this in production, yet.


## Experiments

See and run `main.go` to know how to use this package.

```bash
$ go run main.go

=== Epsilon Greedy Algorithm ===
True mean of bandit(No. 1) = 0.100000, actual = 0.140396
True mean of bandit(No. 2) = 0.200000, actual = 0.187616
True mean of bandit(No. 3) = 0.300000, actual = 0.297815

=== Optimistic Initial Value Algorithm ===
True mean of bandit(No. 1) = 0.100000, actual = 0.247212
True mean of bandit(No. 2) = 0.200000, actual = 0.258311
True mean of bandit(No. 3) = 0.300000, actual = 0.300267
```
