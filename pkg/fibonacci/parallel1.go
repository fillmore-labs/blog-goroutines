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

// Parallel1 calculates the nth Fibonacci number concurrently  by spawning goroutines to calculate the previous two
// numbers.
//
// This allows calculating larger Fibonacci numbers much faster by leveraging multiple CPU cores.
// The goroutines run independently without blocking each other. It returns once both goroutines have completed.
func Parallel1(n int) int {
	if n < 2 {
		return n
	}

	fc1 := make(chan int)
	go func() { fc1 <- Parallel1(n - 1) }()

	fc2 := make(chan int)
	go func() { fc2 <- Parallel1(n - 2) }()

	return <-fc1 + <-fc2
}
