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

package main_test

import (
	"testing"

	. "fillmore-labs.com/blog/goroutines/cmd/try1"
	"go.uber.org/goleak"
)

func TestRun1(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	defer goleak.VerifyNone(t)

	m := Run1(Count)

	t.Logf("*** Finished %d runs - avg %v, stddev %v\n", Count, m.Avg, m.Dev)
}

func BenchmarkRun1(b *testing.B) {
	_ = Run1(b.N)
}
