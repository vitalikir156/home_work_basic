package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWorkersGood(t *testing.T) {
	tests := []struct {
		name  string
		input int
	}{
		{"1", 5},
		{"2", 1000},
		{"3", 1000000},
		//	{"4", 10000000}, // printf забивает оперативу.
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			v, err := wpool(tc.input)
			require.Equal(t, tc.input, v)
			require.NoError(t, err)
		})
	}
}

func TestWorkersBad(t *testing.T) {
	_, err := wpool(-666)
	require.Error(t, err)
}
