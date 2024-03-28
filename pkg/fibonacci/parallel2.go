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

package fibonacci

// Parallel2 calculates the nth Fibonacci number by spawning a goroutine to concurrently calculate the previous number
// (n-1) while calculating the second preceding number (n-2) sequentially in the main goroutine.
//
// This avoids launching an extra goroutine compared to [Parallel1], improving performance.
func Parallel2(n int) int {
	if n < 2 {
		return n
	}

	fc1 := make(chan int)
	go func() { fc1 <- Parallel2(n - 1) }()

	fn2 := Parallel2(n - 2)

	return <-fc1 + fn2
}
