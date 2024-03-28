package main

import (
	"fmt"
	"time"

	"fillmore-labs.com/blog/goroutines/pkg/gc"
)

func main() {
	runConcurrent()
	runNonConcurrent()
}

func runConcurrent() {
	start := time.Now()

	reachable := gc.ConcurrentMark(gc.MakeRootC(), gc.Lookupc)

	fmt.Printf("Concurrent: Found %d reachable nodes in %v\n", len(reachable), time.Since(start))
}

func runNonConcurrent() {
	start := time.Now()

	reachable := gc.Mark(gc.MakeRootC(), gc.Lookupc)

	fmt.Printf("Non-Concurrent: Found %d reachable nodes in %v\n", len(reachable), time.Since(start))
}
