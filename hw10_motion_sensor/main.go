package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {
	stopper := make(chan bool)
	sensdirtydata := make(chan int, 1)
	sensaverdata := make(chan int, 1)
	go sensor(stopper, sensdirtydata)
	go stoptimer(stopper)
	go sensavearage(sensdirtydata, sensaverdata)
	for {
		select {
		case value, ok := <-sensaverdata:
			if !ok {
				fmt.Println("stop")
				return
			}
			fmt.Printf("val %v\n", value)
		default:
			time.Sleep(30 * time.Millisecond)
		}
	}
}

func sensavearage(dd <-chan int, out chan<- int) {
	defer close(out)
	i := 0
	s := 0
	for task := range dd {
		i++
		s += task
		//		fmt.Printf("i %v,task  %v, s %v\n", i, task, s)
		if i > 9 {
			out <- s / 10
			s = 0
			i = 0
		}
	}
}

func sensor(stop <-chan bool, out chan<- int) {
	defer close(out)
	for {
		select {
		case <-stop:
			return
		default:
			s, _ := rand.Int(rand.Reader, big.NewInt(10))
			r := int(s.Int64())
			time.Sleep(60 * time.Millisecond)
			out <- r
		}
	}
}

func stoptimer(stop chan<- bool) {
	time.Sleep(time.Minute)
	close(stop)
}
