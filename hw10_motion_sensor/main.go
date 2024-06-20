package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {
	sensdirtydata := sensor(60)
	sensaverdata := sensavearage(sensdirtydata)
	for value := range sensaverdata {
		fmt.Printf("val %v\n", value)
	}
}

func sensavearage(dd <-chan int) chan int {
	out := make(chan int, 1)
	go func() {
		defer close(out)
		i := 0
		s := 0
		for task := range dd {
			i++
			s += task
			if i > 9 {
				out <- s / 10
				s = 0
				i = 0
			}
		}
		if i > 0 {
			out <- s / i
		}
	}()
	return out
}

func sensor(timer int) chan int {
	to := time.After(time.Duration(timer) * time.Second)
	out := make(chan int, 1)
	go func() {
		defer close(out)
		for {
			select {
			case <-to:
				return
			default:
				s, _ := rand.Int(rand.Reader, big.NewInt(10))
				r := int(s.Int64())
				time.Sleep(60 * time.Millisecond)
				out <- r
			}
		}
	}()
	return out
}
