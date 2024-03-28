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

package gc

type Node string

func Lookupc(ref Node, fn func(Node)) {
	if len(ref) > 8 {
		return
	}

	for _, n := range "ABCD" {
		fn(ref + Node(n))
	}
}

func MakeRootC() <-chan Node {
	rootC := make(chan Node)
	go func() {
		rootC <- ""
		close(rootC)
	}()

	return rootC
}
