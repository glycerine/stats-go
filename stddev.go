/* Copyright 2024 Jason E. Aten, Ph.D. All rights reserved.
// Same LICENSE as Go:

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

   * Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.
   * Redistributions in binary form must reproduce the above
copyright notice, this list of conditions and the following disclaimer
in the documentation and/or other materials provided with the
distribution.
   * Neither the name of Google Inc. nor the names of its
contributors may be used to endorse or promote products derived from
this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package stats

import (
	"math"
)

// StdDevTracker tracks the running stddev and mean
// as each observation is added with AddObs(). Use
// the methods Mean() and SampleStdDev() to retreive
// the summary statistics.
//
// Algorithm reference:
// SANDIA REPORT SAND2008-6212
// Unlimited Release
// Printed September 2008
// "Formulas for Robust, One-Pass Parallel
// Computation of Covariances and Arbitrary-Order Statistical Moments"
// by Philippe Pebay
// http://prod.sandia.gov/techlib/access-control.cgi/2008/086212.pdf (now stale?)
// backup urls:
// https://www.osti.gov/biblio/1028931
// https://www.osti.gov/servlets/purl/1028931
type StdDevTracker struct {

	// W is the sum of all weights seen.
	W float64

	// A is the weighted mean
	A float64

	// Q is the weighted numerator for the variance (the Quadratic term)
	Q float64
}

// Mean return the weighted mean.
func (s *StdDevTracker) Mean() float64 {
	return s.A
}

// SampleStdDev return the weighted sample standard deviation.
func (s *StdDevTracker) SampleStdDev() float64 {
	wvar := s.Q / s.W
	return math.Sqrt(wvar * s.W / (s.W - 1))
}

func (s *StdDevTracker) AddObs(x float64, weight float64) {

	// W is the sum of all weights seen.
	s.W += weight

	// need to save the old value for the updates below.
	a0i := s.A

	// A is the weighted mean
	s.A = a0i + weight*(x-a0i)/s.W

	// update the quadratic term.
	// Q/W gives the weighted variance, when needed.
	s.Q += weight * (x - a0i) * (x - s.A)
}

// MeanAndSampleStdDev returns the mean and sample standard
// deviation using the observations in x.
func MeanAndSampleStdDev(x []float64) (mean, stddev float64) {
	var sdt StdDevTracker
	for _, v := range x {
		sdt.AddObs(v, 1)
	}
	return sdt.Mean(), sdt.SampleStdDev()
}
