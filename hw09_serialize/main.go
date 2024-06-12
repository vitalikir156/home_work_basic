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
type Marshaller interface {
	Marshal() []byte
}
type Unmarshaller interface {
	Unmarshal() Book
}

func (b Book) Marshal() []byte {
	e, _ := json.Marshal(b)
	return e
}

func Unmarshal(b []byte) Book {
	var data Book
	json.Unmarshal(b, &data)
	return data
}

func Slicemarshall(b []Book) []byte {
	s := make([]byte, 0)
	for _, book := range b {
		u := Marshaller.Marshal(book)
		s = append(s, u...)
	}
	return s
}

func main() {
	b := Book{1, "title1", "author", 1991, 12, 4}
	slb := []Book{
		{1, "title1", "author", 1991, 12, 4},
		{2, "title2", "author2", 1992, 22, 3},
	}

	s := Marshaller.Marshal(b)
	// fmt.Println(s)
	fmt.Printf(" %s\n", s)
	slbum := Slicemarshall(slb)
	fmt.Printf("slicem %s\n", slbum)
	e := Unmarshal(s)
	fmt.Println(e)

	pbook := &pb.Book{
		ID: 3,

		Title:  "title1",
		Author: "author",
		Year:   1991,
		Size:   12,
		Rate:   4,
	}
	pbook.ProtoMessage()
	data, _ := proto.Marshal(pbook)
	fmt.Printf(" %s\n", data)
	pbook2 := &pb.Book{}
	proto.Unmarshal(data, pbook2)
	fmt.Println(pbook2)
}
