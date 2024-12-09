# stats-go

Context: https://github.com/golang/go/issues/69264

Comment:

A data analyst typically wants both the mean and
sample standard deviation for a given set of
observations.

For efficiency, we demonstrate here how
to calculate these both with a single pass
over the data.
