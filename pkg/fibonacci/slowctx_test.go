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
	"context"
	"testing"

	. "fillmore-labs.com/blog/goroutines/pkg/fibonacci"
)

func TestSlowCtx(t *testing.T) {
	t.Parallel()

	contextCanceled := func() context.Context {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		return ctx
	}

	type args struct {
		ctx context.Context
		i   int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "27", args: args{ctx: context.Background(), i: 27}, want: 196_418, wantErr: false},
		{name: "canceled", args: args{ctx: contextCanceled(), i: 27}, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := SlowCtx(tt.args.ctx, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlowCtx() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if got != tt.want {
				t.Errorf("SlowCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
