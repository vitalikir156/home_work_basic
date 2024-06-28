package main

import (
	"fmt"
	"sync"
)

func main() {
	wpool()
}

func wpool() {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	v := 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			mu.Lock()
			v++
			mu.Unlock()
			wg.Done()
			fmt.Printf("worker %v: job is done\n", i)
		}(i)
	}
	wg.Wait()
	fmt.Println(v)
}
