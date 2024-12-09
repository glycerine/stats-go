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
	"testing"
)

func TestStdDevTracker(t *testing.T) {

	data := []float64{0.8132389853750429109525, 1.065861676133342594852, 0.4244678384669774429128, -1.739035729372427674377, -0.06181969211179912104814, -0.3944856059607272369405, -1.349434274156593893679, 0.07474778053560837687286, 2.089007542068220146803, 1.441221012264471523778, -0.768796979529236312878, 1.055432475820423876556, -0.3786303606682651645698, 0.03711473417454794893056, 0.04430729462047075517539, 0.6293601401622032076588, -0.6523533146020070727644, -0.7132828283969454563618, -0.9956358346589565533336, 1.452074787914823916779}

	// expect, from R:
	const expectedMean = 0.1036679824039587055617
	const expectedSD = 1.012965904917213721959

	mean, sd := MeanAndSampleStdDev(data)
	if math.Abs(mean-expectedMean) > 1e-8 {
		t.Errorf("observed mean = %g, want %g", mean, expectedMean)
	}
	if math.Abs(sd-expectedSD) > 1e-8 {
		t.Errorf("observed sd = %g, want %g", sd, expectedSD)
	}
}
