package main

import (
	"encoding/json"
	"fmt"

	pb "github.com/vitalikir156/home_work_basic/hw09_serialize/book"
	"google.golang.org/protobuf/proto"
)

type Book struct {
	ID     int64   `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   int64   `json:"year"`
	Size   float32 `json:"size"`
	Rate   float32 `json:"rate"`
}


type Marshaler interface {
    MarshalJSON() ([]byte, error)
}


func (b Book) Marshal() ([]byte, error) {
	e, err := json.Marshal(b)
	return e, err
}

func (b *Book) Unmarshal(data []byte) error {
	err := json.Unmarshal(data, &b)
	return err
}

func MarshalSlice(b []Book) ([]byte, error) {
	e, err := json.Marshal(b)
	return e, err
}

func UnmarshalSlice(data []byte) ([]Book, error) {
	var book []Book
	err := json.Unmarshal(data, &book)
	return book, err
}

func main() {
	b := Book{1, "title1", "author", 1991, 12, 4}
	slb := []Book{
		{1, "title1", "author", 1991, 12, 4},
		{2, "title2", "author2", 1992, 22, 3},
	}
	fmt.Println(slb)
	s, _ := Book.Marshal(b)
	// fmt.Println(s)
	fmt.Printf("s %s\n", s)
	var e Book
	e.Unmarshal(s)
	fmt.Printf("e.unmarsh %v\n", e)

	slm, _ := MarshalSlice(slb)
	fmt.Printf("slicemar %s\n", slm)
	slu, _ := UnmarshalSlice(slm)
	fmt.Printf("sliceunmar %v\n", slu)
slpbook:=&pb.Books{Books: []*pb.Book{{
	Id:     3,
	Title:  "title1",
	Author: "author",
	Year:   1991,
	Size:   12,
	Rate:   4,},
	{
		Id:     4,
		Title:  "title4",
		Author: "author5",
		Year:   19916,
		Size:   123,
		Rate:   1,},
}}
	pbook := &pb.Book{
		Id:     3,
		Title:  "title1",
		Author: "author",
		Year:   1991,
		Size:   12,
		Rate:   4,
	}
	slpbook2, _:=proto.Marshal(slpbook)
slpbook3:=&pb.Books{Books: []*pb.Book{{}}}
	proto.Unmarshal(slpbook2, slpbook3)

	fmt.Printf("asasas %s\n" ,slpbook)
	fmt.Printf("slpbook2 %s\n" ,slpbook2)
	fmt.Printf("slpbook3 %s\n" ,slpbook3)
	pbook.ProtoMessage()
	data, _ := proto.Marshal(pbook)
	// fmt.Printf(" %s\n", data)
	pbook2 := &pb.Book{}
	proto.Unmarshal(data, pbook2)
	fmt.Println(pbook2)
}
