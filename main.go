package main

import (
	"fmt"
	"time"

	"github.com/khekrn/workspace/read"
)

func main() {
	start := time.Now()

	read.ReadChunksInParallel("dummy.txt")

	elapsed := time.Since(start)
	fmt.Printf("\n Total Time = %s", elapsed)
}
