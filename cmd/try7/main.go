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

package main

import (
	"fmt"
	"runtime"
	"time"

	"fillmore-labs.com/blog/goroutines/pkg/fibonacci"
	"fillmore-labs.com/blog/goroutines/pkg/mean"
)

const (
	Count    = 1_000
	sequence = 27
)

func main() {
	start := time.Now()

	m := Run7(Count)

	fmt.Printf("*** Finished %d runs in %v - avg %v, stddev %v\n", Count, time.Since(start), m.Avg, m.Dev)
}

func Run7(c int) mean.Mean {
	numCPU := runtime.GOMAXPROCS(0)

	loopCount := c / numCPU
	remainder := c - (loopCount * numCPU)

	var g int
	aggregates := make(chan mean.Aggregate)
	for i := range numCPU {
		count := loopCount
		if i == 0 {
			count += remainder
		}

		g++
		go func() {
			var a mean.Aggregate

			for range count {
				queryStart := time.Now()
				_ = fibonacci.Slow(sequence)
				a.Update(float64(time.Since(queryStart)))
			}

			aggregates <- a
		}()
	}

	var s mean.Aggregate
	for range g {
		s = s.Combine(<-aggregates)
	}

	return s.Finalize()
}
