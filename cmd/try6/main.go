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
	"context"
	"fmt"
	"runtime"
	"sync/atomic"
	"time"

	"fillmore-labs.com/blog/goroutines/pkg/fibonacci"
	"fillmore-labs.com/blog/goroutines/pkg/mean"
	"golang.org/x/sync/semaphore"
)

const (
	Count    = 1_000
	sequence = 27
	Timeout  = 100 * time.Millisecond
)

func main() {
	start := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	d, f, m := Run6(ctx, Count)

	fmt.Printf("*** Finished %d runs (%d failed) in %v - avg %v, stddev %v\n", d, f, time.Since(start), m.Avg, m.Dev)
}

func Run6(ctx context.Context, c int) (d, f int64, r mean.Mean) {
	s := mean.New()
	var done, failed atomic.Int64

	numCPU := int64(runtime.GOMAXPROCS(0))
	pool := semaphore.NewWeighted(numCPU)
	for range c {
		queryStart := time.Now()

		if err := pool.Acquire(ctx, 1); err != nil {
			break
		}
		go func() {
			defer pool.Release(1)

			_, err := fibonacci.SlowCtx(ctx, sequence)

			if err == nil {
				s.Add(time.Since(queryStart))

				done.Add(1)
			} else {
				failed.Add(1)
			}
		}()
	}
	_ = pool.Acquire(context.WithoutCancel(ctx), numCPU)

	return done.Load(), failed.Load(), s.Result()
}
