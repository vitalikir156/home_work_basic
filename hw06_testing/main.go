package main

import (
	"fmt"

	"github.com/vitalikir156/home_work_basic/hw06_testing/chessboard"
	"github.com/vitalikir156/home_work_basic/hw06_testing/fixapp"
	"github.com/vitalikir156/home_work_basic/hw06_testing/shapes"
	"github.com/vitalikir156/home_work_basic/hw06_testing/structcompar"
)
type Sts string
func (t Sts) String() string {
	
	return string(t)+"sdsd\nsdsd"
}
func main() {
a:=Sts("sdf")
fmt.Println(a)
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
	_ = chessboard.Auto(8, true)
	// structcompar.Structcompar()
	var book3 structcompar.Book
	book3.SetYear(2001)
	fmt.Println(book3.Year())
	fmt.Println(a)
}
