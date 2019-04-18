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
True mean of bandit(No. 1) = 0.100000, actual = -0.710439
True mean of bandit(No. 2) = 0.200000, actual = 0.012979
True mean of bandit(No. 3) = 0.300000, actual = 0.275188

=== Optimistic Initial Value Algorithm ===
Time 0.000018 [sec] for 100-iterations
True mean of bandit(No. 1) = 0.100000, actual = 0.435606
True mean of bandit(No. 2) = 0.200000, actual = 0.469415
True mean of bandit(No. 3) = 0.300000, actual = 0.441793

=== UCB1 Algorithm ===
Time 0.000031 [sec] for 100-iterations
True mean of bandit(No. 1) = 0.100000, actual = -0.250816
True mean of bandit(No. 2) = 0.200000, actual = 0.378808
True mean of bandit(No. 3) = 0.300000, actual = 0.164395

=== Epsilon Greedy Algorithm ===
Time 0.001542 [sec] for 10000-iterations
True mean of bandit(No. 1) = 0.100000, actual = 0.107990
True mean of bandit(No. 2) = 0.200000, actual = 0.268195
True mean of bandit(No. 3) = 0.300000, actual = 0.299267

=== Optimistic Initial Value Algorithm ===
Time 0.001393 [sec] for 10000-iterations
True mean of bandit(No. 1) = 0.100000, actual = 0.247309
True mean of bandit(No. 2) = 0.200000, actual = 0.249943
True mean of bandit(No. 3) = 0.300000, actual = 0.318681

=== UCB1 Algorithm ===
Time 0.002630 [sec] for 10000-iterations
True mean of bandit(No. 1) = 0.100000, actual = 0.152315
True mean of bandit(No. 2) = 0.200000, actual = 0.176517
True mean of bandit(No. 3) = 0.300000, actual = 0.296878

=== Epsilon Greedy Algorithm ===
Time 0.146866 [sec] for 1000000-iterations
True mean of bandit(No. 1) = 0.100000, actual = 0.099697
True mean of bandit(No. 2) = 0.200000, actual = 0.201168
True mean of bandit(No. 3) = 0.300000, actual = 0.300164

=== Optimistic Initial Value Algorithm ===
Time 0.107016 [sec] for 1000000-iterations
True mean of bandit(No. 1) = 0.100000, actual = 0.193490
True mean of bandit(No. 2) = 0.200000, actual = 0.194106
True mean of bandit(No. 3) = 0.300000, actual = 0.300441

=== UCB1 Algorithm ===
Time 0.212366 [sec] for 1000000-iterations
True mean of bandit(No. 1) = 0.100000, actual = 0.141993
True mean of bandit(No. 2) = 0.200000, actual = 0.206228
True mean of bandit(No. 3) = 0.300000, actual = 0.301404
```
