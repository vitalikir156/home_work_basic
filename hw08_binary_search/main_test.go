package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinSearchGood(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"0", 1, 0},
		{"1", 5, 1},
		{"2", 7, 2},
		{"3", 9, 3},
		{"4", 13, 4},
		{"5", 57, 5},
		{"6", 58, 6},
		{"7", 61, 7},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got, err := BinSearch([]int{1, 5, 7, 9, 13, 57, 58, 61}, tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, got)
		})
	}
}

func TestBinSearchBad(t *testing.T) {
	tests := []struct {
		name  string
		input int
	}{
		{"bad0", 2},
		{"bad1", 11},
		{"bad2", 98},
		{"bad3", 0},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			_, err := BinSearch([]int{1, 5, 7, 9, 13, 57, 58, 61}, tc.input)
			require.Error(t, err)
		})
	}
}
