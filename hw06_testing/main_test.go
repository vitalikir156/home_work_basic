package main_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vitalikir156/home_work_basic/hw06_testing/chessboard"
	"github.com/vitalikir156/home_work_basic/hw06_testing/fixapp"
	"github.com/vitalikir156/home_work_basic/hw06_testing/fixapp/types"
	"github.com/vitalikir156/home_work_basic/hw06_testing/shapes"
	"github.com/vitalikir156/home_work_basic/hw06_testing/structcompar"
)

func TestChessBoardGood(t *testing.T) {
	tests := []struct {
		name    string
		howmuch int
		gm      bool
		reqire  string
	}{
		{"8x8, no GM", 8, false, "# # # # \n # # # #\n# # # # \n # # # #\n# # # # \n # # # #\n# # # # \n # # # #\n"},
		{"3x3, with GM", 3, true, "#   # \n  #   \n#   # \n"},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got, err := chessboard.Auto(tc.howmuch, tc.gm)
			require.NoError(t, err)
			require.Equal(t, tc.reqire, got)
		})
	}
}

func TestChessBoardBad(t *testing.T) {
	_, err := chessboard.Auto(-7, false)
	require.Error(t, err)
}

func TestFixAppGood(t *testing.T) {
	staff, err := fixapp.Fixapp("fixapp/data.json")
	req := []types.Employee{
		{UserID: 10, Age: 25, Name: "Rob", DepartmentID: 3},
		{UserID: 11, Age: 30, Name: "George", DepartmentID: 2},
	}
	require.NoError(t, err)
	require.Equal(t, staff, req)
}

func TestFixAppBadPatch(t *testing.T) {
	_, err := fixapp.Fixapp("badpatch/data.json")
	require.Error(t, err)
}

func TestFixAppBadJSON(t *testing.T) {
	_, err := fixapp.Fixapp("fixapp/baddata.json")
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

func TestStructCompGood(t *testing.T) {
	tests := []struct {
		name          string
		book1year     int
		book2year     int
		book1size     float32
		book2size     float32
		book1rate     float32
		book2rate     float32
		modeofcompare string
		expected      bool
	}{
		{"testone- year", 1974, 2010, 47.91, 291.1, 3.3, 4.8, "year", false},
		{"two- size", 1974, 2010, 47.91, 291.1, 3.3, 4.8, "size", false},
		{"three- rate", 2016, 2010, 371.67, 291.1, 4.8, 4.6, "rate", true},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			var s1 structcompar.Book

			s1.SetRate(tc.book1rate)
			s1.SetSize(tc.book1size)
			s1.SetYear(tc.book1year)

			var s2 structcompar.Book

			s2.SetRate(tc.book2rate)
			s2.SetSize(tc.book2size)
			s2.SetYear(tc.book2year)

			got, err := structcompar.Structcompar(s1, s2, tc.modeofcompare)
			require.NoError(t, err)
			require.Equal(t, tc.expected, got)
		})
	}
}

func TestStructCompBad(t *testing.T) {
	var s1 structcompar.Book
	var s2 structcompar.Book

	_, err := structcompar.Structcompar(s1, s2, "badmode")
	require.Error(t, err)
}
