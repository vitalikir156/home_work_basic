package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAvearage(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"1", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5},
		{"2", []int{9, 9, 9, 9, 9}, 9},
		{"3", []int{4, 3, 2, 1, 2, 3, 4, 3, 1, 2}, 2},
		{"4", []int{9, 8, 9, 7, 6, 5, 7, 9, 9, 6}, 7},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			sensdirtydata := make(chan int, 1)
			sensaverdata := sensavearage(sensdirtydata)
			for _, v := range tc.input {
				sensdirtydata <- v
			}
			close(sensdirtydata)
			require.Equal(t, tc.expected, <-sensaverdata)
		})
	}
}

func TestSensor(t *testing.T) {
	ta := time.After(time.Minute)

	sensdirtydata := sensor()
	var err error
	go func() {
		for task := range sensdirtydata {
			if 10 < task || task < 0 {
				err = fmt.Errorf("out of range: %v", task)
			}
		}
	}()

	<-ta // wait minute
	_, ok := <-sensdirtydata
	if ok {
		err = fmt.Errorf("timedout but channel is open")
	}

	require.NoError(t, err)
}
