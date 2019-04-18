# gorl

gorl is a go language implementation of reinforcement learning.

This package is under development.
You shouldn't use this in production, yet.


## Experiments

See and run `main.go` to know how to use this package.

```bash
$ go run main.go

=== Epsilon Greedy Algorithm ===
Time 0.000028 [sec] for 100-iterations
True mean of bandit(No. 1) = 0.100000, actual = -0.198847
True mean of bandit(No. 2) = 0.200000, actual = 0.371966
True mean of bandit(No. 3) = 0.300000, actual = 0.259781

=== Optimistic Initial Value Algorithm ===
Time 0.000017 [sec] for 100-iterations
True mean of bandit(No. 1) = 0.100000, actual = 0.307894
True mean of bandit(No. 2) = 0.200000, actual = 0.312110
True mean of bandit(No. 3) = 0.300000, actual = 0.389637

=== UCB1 Algorithm ===
Time 0.000027 [sec] for 100-iterations
True mean of bandit(No. 1) = 0.100000, actual = -0.170127
True mean of bandit(No. 2) = 0.200000, actual = 0.126382
True mean of bandit(No. 3) = 0.300000, actual = 0.514704

=== Epsilon Greedy Algorithm ===
Time 0.001515 [sec] for 10000-iterations
True mean of bandit(No. 1) = 0.100000, actual = 0.109374
True mean of bandit(No. 2) = 0.200000, actual = 0.196008
True mean of bandit(No. 3) = 0.300000, actual = 0.296269

=== Optimistic Initial Value Algorithm ===
Time 0.001195 [sec] for 10000-iterations
True mean of bandit(No. 1) = 0.100000, actual = 0.146754
True mean of bandit(No. 2) = 0.200000, actual = 0.199270
True mean of bandit(No. 3) = 0.300000, actual = 0.123285

=== UCB1 Algorithm ===
Time 0.002421 [sec] for 10000-iterations
True mean of bandit(No. 1) = 0.100000, actual = 0.058483
True mean of bandit(No. 2) = 0.200000, actual = 0.066505
True mean of bandit(No. 3) = 0.300000, actual = 0.304352

=== Epsilon Greedy Algorithm ===
Time 0.151637 [sec] for 1000000-iterations
True mean of bandit(No. 1) = 0.100000, actual = 0.103865
True mean of bandit(No. 2) = 0.200000, actual = 0.199959
True mean of bandit(No. 3) = 0.300000, actual = 0.296444

=== Optimistic Initial Value Algorithm ===
Time 0.107858 [sec] for 1000000-iterations
True mean of bandit(No. 1) = 0.100000, actual = 0.276607
True mean of bandit(No. 2) = 0.200000, actual = 0.264976
True mean of bandit(No. 3) = 0.300000, actual = 0.299464

=== UCB1 Algorithm ===
Time 0.226486 [sec] for 1000000-iterations
True mean of bandit(No. 1) = 0.100000, actual = 0.057508
True mean of bandit(No. 2) = 0.200000, actual = 0.173537
True mean of bandit(No. 3) = 0.300000, actual = 0.299723
```
