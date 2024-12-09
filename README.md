# stats-go

Context: https://github.com/golang/go/issues/69264

Comment:

A data analyst typically wants both the mean and
sample standard deviation for a given set of
observations.

For efficiency, we demonstrate here how
to calculate these both with a single pass
over the data. We make only one call to:

~~~
// MeanAndSampleStdDev returns the mean and sample standard
// deviation using the observations in x.
func MeanAndSampleStdDev(x []float64) (mean, stddev float64)
	var sdt StdDevTracker
	for _, v := range x {
		sdt.AddObs(v, 1)
	}
	return sdt.Mean(), sdt.SampleStdDev()
}
~~~


