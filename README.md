# gorl

gorl is a go language implementation of reinforcement learning.

This package is under development.
You shouldn't use this in production, yet.


## Experiments

See and run `main.go` to know how to use this package.

```bash
$ go run main.go

=== Epsilon Greedy Algorithm ===
True mean of bandit(No. 1) = 0.100000, actual = 0.112452
True mean of bandit(No. 2) = 0.200000, actual = 0.221077
True mean of bandit(No. 3) = 0.300000, actual = 0.308172

=== Optimistic Initial Value Algorithm ===
True mean of bandit(No. 1) = 0.100000, actual = 0.287140
True mean of bandit(No. 2) = 0.200000, actual = 0.281862
True mean of bandit(No. 3) = 0.300000, actual = 0.295689

=== UCB1 Algorithm ===
True mean of bandit(No. 1) = 0.100000, actual = 0.067241
True mean of bandit(No. 2) = 0.200000, actual = 0.152241
True mean of bandit(No. 3) = 0.300000, actual = 0.292240
```
