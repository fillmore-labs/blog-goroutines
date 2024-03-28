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
	"math"
	"time"
)

// Mean represents the average and standard deviation of a set of durations.
type Mean struct {
	Avg time.Duration
	Dev time.Duration
}

// Aggregate accumulates the mean and standard deviation of a set of durations.
type Aggregate struct {
	count    int
	mean, m2 float64
}

// Update adds a new value to the aggregate.
func (a *Aggregate) Update(value float64) {
	a.count++
	delta := value - a.mean
	a.mean += delta / float64(a.count)
	delta2 := value - a.mean
	a.m2 += delta * delta2
}

// Finalize returns the final mean and standard deviation of the values added to the aggregate.
// It is called after all values have been added.
func (a *Aggregate) Finalize() Mean {
	var stdDev float64
	if a.count > 0 {
		stdDev = math.Sqrt(a.m2 / float64(a.count))
	}

	return Mean{Avg: time.Duration(a.mean), Dev: time.Duration(stdDev)}
}

// Combine combines two aggregates.
func (a *Aggregate) Combine(b Aggregate) Aggregate {
	count := a.count + b.count
	mean := (float64(a.count)*a.mean + float64(b.count)*b.mean) / float64(count)
	delta := b.mean - a.mean
	m2 := a.m2 + b.m2 + delta*delta*float64(a.count)*float64(b.count)/(float64(count))

	return Aggregate{count: count, mean: mean, m2: m2}
}
