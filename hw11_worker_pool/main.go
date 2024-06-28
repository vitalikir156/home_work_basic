package main

import (
	"fmt"
	"sync"
)

func main() {
	v, err := wpool(10)
	fmt.Println(v, err)
}

func wpool(nor int) (int, error) {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	v := 0
	if nor < 1 {
		return 0, fmt.Errorf("bad number of goroutines: %v", nor)
	}
	for i := 0; i < nor; i++ {
		wg.Add(1)
		go func(i int) {
			mu.Lock()
			v++
			fmt.Printf("worker %v: job is done\n", i) // если не выделить в mutex то забивает сокет раньше времени
			mu.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()

	return v, nil
}
