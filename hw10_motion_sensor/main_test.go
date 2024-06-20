package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAvearage(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"1", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5},
		{"2", []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}, 9},
		{"3", []int{4, 3, 2, 1, 2, 3, 4, 3, 1, 2}, 2},
		{"4", []int{9, 8, 9, 7, 6, 5, 7, 9, 9, 6}, 7},
	}
	for _, tc := range tests {
		sensdirtydata := make(chan int, 1)
		sensaverdata := make(chan int, 1)
		go sensavearage(sensdirtydata, sensaverdata)
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			for _, v := range tc.input {
				sensdirtydata <- v
			}

			require.Equal(t, tc.expected, <-sensaverdata)
		})
		close(sensdirtydata)
	}
}
