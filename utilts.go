package main

import (
	"fmt"
	"time"
)

func measureTime(f func() []int, label string) {
	start := time.Now()
	result := f()
	elapsed := time.Since(start)

	fmt.Printf("Algorithm: %s, Result: %v, Time: %s\n", label, result, elapsed)
}
