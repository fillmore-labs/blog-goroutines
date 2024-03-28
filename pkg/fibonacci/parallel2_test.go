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

package fibonacci_test

import (
	"testing"

	. "fillmore-labs.com/blog/goroutines/pkg/fibonacci"
)

func TestParallel2(t *testing.T) {
	t.Parallel()

	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "27", args: args{i: 27}, want: 196_418},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := Parallel2(tt.args.i); got != tt.want {
				t.Errorf("Parallel2() = %v, want %v", got, tt.want)
			}
		})
	}
}