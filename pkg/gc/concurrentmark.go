// Copyright The containerd Authors.
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

import (
	"sync"
)

func ConcurrentMark(root <-chan Node, refs func(Node, func(Node))) map[Node]struct{} {
	var (
		grays = make(chan Node)
		seen  = map[Node]struct{}{} // or not "white", basically "seen"
		wg    sync.WaitGroup
	)

	go func() {
		for gray := range grays {
			if _, ok := seen[gray]; ok {
				wg.Done()

				continue
			}
			seen[gray] = struct{}{} // post-mark this as non-white

			go func() {
				defer wg.Done()

				send := func(n Node) {
					wg.Add(1)
					grays <- n
				}

				refs(gray, send)
			}()
		}
	}()

	for r := range root {
		wg.Add(1)
		grays <- r
	}

	// Wait for outstanding grays to be processed
	wg.Wait()

	close(grays)

	return seen
}
