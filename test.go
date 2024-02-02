package main

import (
	"fmt"
	"sync"
)

func main() {
	values := []int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup

	for _, val := range values {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			fmt.Println(v)
		}(val)
	}

	wg.Wait()
	fmt.Println("hello world")
}
