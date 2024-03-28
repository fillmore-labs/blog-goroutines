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

	m := Run1(Count)

	fmt.Printf("*** Finished %d runs in %v - avg %v, stddev %v\n", Count, time.Since(start), m.Avg, m.Dev)
}

func Run1(c int) mean.Mean {
	s := mean.New()

	for range c {
		queryStart := time.Now()

		_ = fibonacci.Slow(sequence)

		s.Add(time.Since(queryStart))
	}

	return s.Result()
}
