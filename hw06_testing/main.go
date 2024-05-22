package main

import (
	"fmt"

	"github.com/vitalikir156/home_work_basic/hw06_testing/chessboard"
	"github.com/vitalikir156/home_work_basic/hw06_testing/fixapp"
	"github.com/vitalikir156/home_work_basic/hw06_testing/shapes"
	"github.com/vitalikir156/home_work_basic/hw06_testing/structcompar"
)

func main() {
	strcmpr()
	fixapptst()
	shapestst()
}

func strcmpr() {
	var s1 structcompar.Book
	s1.SetID(123)
	s1.SetAuthor("author one")
	s1.SetRate(2)
	s1.SetSize(231)
	s1.SetYear(2001)
	s1.SetTitle("FoxOne")
	var s2 structcompar.Book
	s2.SetID(1763)
	s2.SetAuthor("author two")
	s2.SetRate(4)
	s2.SetSize(12)
	s2.SetYear(1937)
	s2.SetTitle("Booktwo")
	comp, err := structcompar.Structcompar(s1, s2, "size")
	fmt.Println(comp)
	fmt.Println(err)
}

func fixapptst() {
	staff, err := fixapp.Fixapp("fixapp/data.json")
	fmt.Println(err)
	_, err = fixapp.Fixapp("fixapp/baddata.json")

	fmt.Println(err)
	_, err = fixapp.Fixapp("fixapp/badddata.json")

	fmt.Println(err)
	fmt.Println(staff)
}

func shapestst() {
	a, _ := shapes.CalculateArea(shapes.Rectangle{Width: 9, Height: 17})
	_, _ = chessboard.Auto(8, true)
	fmt.Println(a)
}
