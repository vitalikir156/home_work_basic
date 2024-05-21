package main_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vitalikir156/home_work_basic/hw06_testing/chessboard"
	"github.com/vitalikir156/home_work_basic/hw06_testing/fixapp"
	"github.com/vitalikir156/home_work_basic/hw06_testing/shapes"
)

func TestChessBoardGood(t *testing.T) {
	err := chessboard.Auto(8, false)
	require.NoError(t, err)
}

func TestChessBoardBad(t *testing.T) {
	err := chessboard.Auto(-7, false)
	require.Error(t, err)
}

func TestFixAppGood(t *testing.T) {
	err := fixapp.Fixapp("fixapp/data.json")
	require.NoError(t, err)
}

func TestFixAppBad(t *testing.T) {
	err := fixapp.Fixapp("badpatch/data.json")
	require.Error(t, err)
}

func TestShapesGood(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected float64
	}{
		{"circle", shapes.Circle{Radius: 6}, 113.09733552923255},
		{"rectangle", shapes.Rectangle{Width: 9, Height: 17}, 153},
		{"triangle", shapes.Triangle{Base: 51, Height: 3}, 76.5},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got, err := shapes.CalculateArea(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, got)
		})
	}
}

func TestShapesBad(t *testing.T) {
	_, err := shapes.CalculateArea(967)
	require.Error(t, err)
}
