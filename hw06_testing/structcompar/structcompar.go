package structcompar

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

const (
	year = iota
	size
	rate
)

func NewComparator(ftc int) *Comparator {
	return &Comparator{ftc}
}

func (c Comparator) Compare(bookOne, bookTwo *Book) bool {
	switch c.fieldComapre {
	case year:
		{
			return bookOne.year > bookTwo.year
		}
	case size:
		{
			return bookOne.size > bookTwo.size
		}
	case rate:
		{
			return bookOne.rate > bookTwo.rate
		}
	}
	return false
}

func Structcompar(book1, book2 Book) {
	book1 = Book{2134, "bookone", "author one", 2005, 63.2, 5}
	book2 = Book{6568, "booktwo", "author two", 11974, 4099.01, 4.7}
	fmt.Println(Comparator.Compare(*NewComparator(year), &book1, &book2))
}

func (p Book) ID() int {
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

func (p *Book) SetID(id int) {
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
