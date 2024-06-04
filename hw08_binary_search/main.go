package main

import "fmt"

var sl []int

func BinSearch(in []int, srch int) (int, error) {
	max := len(in)
	min := 0

	for min < max {
		half := (max - min) / 2
		switch {
		case in[min+half] == srch:
			return (min + half), nil
		case in[min+half] < srch:
			if half < 1 {
				half = 1
			}
			min += half
		case in[min+half] > srch:
			if half < 1 {
				half = 1
			}
			max -= half
		}
	}
	return 0, fmt.Errorf("integer not found")
}

func main() {
	sl = []int{1, 4, 8, 9, 12, 15, 19}

	bs, _ := BinSearch(sl, 12) // bs not an BullShit, BinarySearch
	fmt.Println(bs)
}
