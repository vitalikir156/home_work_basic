package main

import (
	"fmt"

	"github.com/vitalikir156/home_work_basic/hw06_testing/chessboard"
	"github.com/vitalikir156/home_work_basic/hw06_testing/fixapp"
	"github.com/vitalikir156/home_work_basic/hw06_testing/shapes"
)

func main() {
	a, _ := shapes.CalculateArea(shapes.Rectangle{Width: 9, Height: 17})
	_ = chessboard.Auto(8, true)
	_ = fixapp.Fixapp("")
	fmt.Println(a)
}
