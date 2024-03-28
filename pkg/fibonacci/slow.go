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

// Slow recursively calculates the nth Fibonacci number.
//
// It implements the mathematical definition of the Fibonacci sequence,
// where each number is the sum of the two numbers before it.
//
// However, because it is calculating the values recursively without caching or goroutines,
// it will be very slow for large numbers due to redundant function calls.
func Slow(n int) int {
	if n < 2 {
		return n
	}

	fn1 := Slow(n - 1)
	fn2 := Slow(n - 2)

	return fn1 + fn2
}
