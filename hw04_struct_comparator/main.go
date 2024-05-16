package main

import (
	"fmt"
)

type Book struct {
	id            int
	title, author string
	year          int
	size, rate    float32
}
type Doublebook struct {
	dbook1, dbook2 Book
}

func main() {
	book1 := Book{}
	book2 := Book{}
	fmt.Println("Enter book 1")
	book1.SetBook()
	fmt.Println("Enter book 2")
	book2.SetBook()
	doublebook := Doublebook{book1, book2}
	fmt.Println(doublebook)
	fmt.Printf("Compare mode set:\n 1 - year \n 2 - size \n 3 - rate \n >")
	var mode int8
	_, err := fmt.Scanln(&mode)
	if err != nil || mode > 3 {
		fmt.Printf("bad compare mode selection: %v! \nExiting!\n", mode)
		return
	}
	comparator := doublebook.Comparebooks(mode)
	fmt.Println(comparator)
}

func (d Doublebook) Comparebooks(mode ...int8) bool {
	fmt.Println(mode)
	compint := 0
	var boolout bool
	for _, v := range mode {
		// mode переключает режимы сравнения, можно получить несколько, выводится первый успешный результат
		fmt.Println(v)
		switch v {
		case 1:
			{
				if d.dbook1.year > d.dbook2.year && compint == 0 {
					compint = 1
				} else if d.dbook1.year < d.dbook2.year && compint == 0 {
					compint = 2
				}
			}
		case 2:
			{
				if d.dbook1.size > d.dbook2.size && compint == 0 {
					compint = 1
				} else if d.dbook1.size < d.dbook2.size && compint == 0 {
					compint = 2
				}
			}
		case 3:
			{
				if d.dbook1.rate > d.dbook2.rate && compint == 0 {
					compint = 1
				} else if d.dbook1.rate < d.dbook2.rate && compint == 0 {
					compint = 2
				}
			}
		}
	}
	switch compint {
	case 1:
		{
			boolout = true
		}
	case 2:
		{
			boolout = false
		}
	default:
		{
			break
		}
	}
	return boolout
}

/*
	func (b Book) Books() (int, string, string, int, float32, float32) {
		return b.id, b.title, b.author, b.year, b.size, b.rate
	}
*/
func (b *Book) SetBook() {
	b.id = intreader("Id")
	b.title = stringreader("title")
	b.author = stringreader("author")
	b.year = intreader("Year")
	b.size = float32reader("Size")
	b.rate = float32reader("Rate")
}

func intreader(name string) int {
	var r int
	fmt.Println("Enter", name, ":")
	_, err := fmt.Scanln(&r)
	if err != nil {
		r = 0
		fmt.Println("Fail! Using default value of", name, ": 0")
	} else {
		fmt.Println(name, "successfully readed:", r)
	}

	return r
}

func stringreader(name string) string {
	var r string
	fmt.Println("Enter", name, ":")
	_, err := fmt.Scanln(&r)
	if err != nil {
		r = ""
		fmt.Println("Fail! Using default value of", name, ": 0")
	} else {
		fmt.Println(name, "successfully readed:", r)
	}

	return r
}

func float32reader(name string) float32 {
	var r float32
	fmt.Println("Enter", name, ":")
	_, err := fmt.Scanln(&r)
	if err != nil {
		r = 0
		fmt.Println("Fail! Using default value of", name, ": 0")
	} else {
		fmt.Println(name, "successfully readed:", r)
	}

	return r
}
