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

func Mark(root <-chan Node, refs func(Node, func(Node))) map[Node]struct{} {
	var grays []Node
	for r := range root {
		grays = append(grays, r)
	}

	seen := map[Node]struct{}{} // or not "white", basically "seen"
	for len(grays) > 0 {
		var next []Node
		for _, gray := range grays {
			if _, ok := seen[gray]; ok {
				continue
			}
			seen[gray] = struct{}{} // post-mark this as non-white

			next = append(next, gray)
		}
		grays = nil

		send := func(n Node) {
			grays = append(grays, n)
		}

		for _, ref := range next {
			refs(ref, send)
		}
	}

	return seen
}
