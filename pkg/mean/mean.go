// Copyright 2024 Oliver Eikemeier. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package mean

import (
	"time"
)

// Summarizer collects durations and returns the calculated [Mean].
type Summarizer struct {
	in  chan<- float64
	out <-chan Mean
}

// New returns a new [Summarizer] and begins collecting durations.
//
// The returned [Summarizer] must be finalized by  calling [Summarizer.Result].
func New() Summarizer {
	in := make(chan float64)
	out := make(chan Mean)
	go variance(in, out)

	return Summarizer{in, out}
}

// Add the given duration to the result.
func (m Summarizer) Add(d time.Duration) {
	m.in <- float64(d.Nanoseconds())
}

// Result returns the calculated [Mean].
func (m Summarizer) Result() Mean {
	close(m.in)

	return <-m.out
}

// variance computes variance using Welford's algorithm.
func variance(in <-chan float64, out chan<- Mean) {
	var a Aggregate

	for d := range in {
		a.Update(d)
	}

	out <- a.Finalize()
	close(out)
}
