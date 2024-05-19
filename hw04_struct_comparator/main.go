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
type Comparator struct {
	fieldComapre int
}

// func NewComparator(year int)*Comparator{}.
func (c Comparator) Compare(bookOne, bookTwo *Book) bool {
	switch c.fieldComapre {
	case 1:
		{
			return bookOne.year > bookTwo.year
		}
	case 2:
		{
			return bookOne.size > bookTwo.size
		}
	case 3:
		{
			return bookOne.rate > bookTwo.rate
		}
	}
	return false
}

func main() {
	book1 := Book{2134, "bookone", "author one", 2005, 63.2, 5}
	book2 := Book{6568, "booktwo", "author two", 1974, 4099.01, 4.7}
	compare := Comparator{1}
	c := Comparator.Compare(compare, &book1, &book2)
	fmt.Println(c)
}

func (p Book) Id() int {
	return p.id
}

func (p Book) Title() string {
	return p.title
}

func (p Book) Author() string {
	return p.author
}

func (p Book) Year() int {
	return p.year
}

func (p Book) Size() float32 {
	return p.size
}

func (p Book) Rate() float32 {
	return p.rate
}

func (p *Book) SetId(id int) {
	p.id = id
}

func (p *Book) SetTitle(title string) {
	p.title = title
}

func (p *Book) SetAuthor(author string) {
	p.author = author
}

func (p *Book) SetYear(year int) {
	p.year = year
}

func (p *Book) SetSize(size float32) {
	p.size = size
}

func (p *Book) SetRate(rate float32) {
	p.rate = rate
}
