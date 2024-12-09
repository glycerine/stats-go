# stats-go

Context: https://github.com/golang/go/issues/69264

Comment:

A data analyst typically wants both the mean and
sample standard deviation for a given set of
observations.

For efficiency, we demonstrate here how
to calculate both these statistics with a single pass
over the data. 

The implementation of StdDevTracker could easily
be inlined into a production version of this function.

~~~
// MeanSd returns the mean and sample standard
// deviation from a single pass through the observations in x.
func MeanSd(x []float64) (mean, stddev float64)
	var sdt StdDevTracker
	for _, v := range x {
		sdt.AddObs(v, 1)
	}
	return sdt.Mean(), sdt.SampleStdDev()
}
~~~


