package shapes

import (
	"errors"
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}
type Circle struct {
	Radius float64
}
type Rectangle struct {
	Width  float64
	Height float64
}
type Triangle struct {
	Base   float64
	Height float64
}

func (a Circle) Area() float64    { return a.Radius * a.Radius * math.Pi }
func (a Rectangle) Area() float64 { return a.Height * a.Width }
func (a Triangle) Area() float64  { return a.Height * a.Base / 2 }

// func (t Shape) calculateArea(s any).
func ShapeMain() {
	c := Circle{5}
	r := Rectangle{10, 5}
	t := Triangle{8, 6}
	shapeareac, err := calculateArea(c)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(shapeareac)
	}
	shapearear, err := calculateArea(r)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(shapearear)
	}
	shapeareat, err := calculateArea(t)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(shapeareat)
	}
	shapearea, err := calculateArea(98989)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(shapearea)
	}
}

func calculateArea(i any) (float64, error) {
	c, ok := i.(Shape)
	if ok {
		return Shape.Area(c), nil
	}
	err := errors.New("the passed object is not a shape")
	return 0, err
}
