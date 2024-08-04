package main

// типа клиент, потыкать серверную часть запросами.
import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

var i bool

func main() {
	for !i {
		Client("http://localhost:9989")
	}
}

func Client(u string) {
	fmt.Print(`
	0 exit
	1 /get 
	2 /get?table=users
	3 /get?table=userorders&srch=a342e4da-1438-4d60-bd4e-7beef2791324
	4 /get?table=usersumm&srch=a342e4da-1438-4d60-bd4e-7beef2791324
	5 /get?table=useraver&srch=a342e4da-1438-4d60-bd4e-7beef2791324
	6 /save", {"desc": "ATTINY 13, 8b, 9.6mhz 1k flash","price": 170}
	7 /save?table=tableupdate", { "id":4, "desc": "ATTINY 13, 8b, 9.6mhz 1k flash","price": 170}
	8 /save?table=deleterow", { "id":9}
	9 /save?table=usercreate", {"name": "Edward Bill","email": "eb@rambler.ru", "password":"1212"}
	10 /save?table=userupdate", {"id": "f89813e7-27b0-478d-b20e-f36a9ba9d759","email": "eb@rambler.ru"}
	11 /save?table=userdelete", {"id": "f89813e7-27b0-478d-b20e-f36a9ba9d759"}
	12 /save?table=orderadd", {"uid": "a30aaed1-7522-4c37-a6a9-b5774fe872e6"}
	13 /save?table=orderfill", {"oid": 14, "pid":2}
	14 /save?table=orderdel", {"oid": "15"}
	select and hit enter:`)
	var s int

	fmt.Scanln(&s)
	fmt.Println(s)
	switch s {
	case 0:
		{
			i = true
		}
	case 1:
		{
			Getter((u + "/get"))
		}
	case 2:
		{
			Getter((u + "/get?table=users"))
		}
	case 3:
		{
			Getter((u + "/get?table=userorders&srch=a30aaed1-7522-4c37-a6a9-b5774fe872e6"))
		}
	case 4:
		{
			Getter((u + "/get?table=usersumm&srch=a342e4da-1438-4d60-bd4e-7beef2791324"))
		}
	case 5:
		{
			Getter((u + "/get?table=useraver&srch=a342e4da-1438-4d60-bd4e-7beef2791324"))
		}
	case 6:
		{
			Pusher(u+"/save", `{"desc": "ATTINY 13, 8b, 9.6mhz 1k flash","price": 170}`)
		}
	case 7:
		{
			Pusher(u+"/save?table=tableupdate", `{ "id":4, "desc": "ATTINY 13, 8b, 9.6mhz 1k flash","price": 170}`)
		}
	case 8:
		{
			Pusher(u+"/save?table=deleterow", `{ "id":9}`)
		}
	case 9:
		{
			Pusher(u+"/save?table=usercreate", `{"name": "Edward Bill","email": "eb@rambler.ru", "password":"1212"}`)
		}
	case 10:
		{
			Pusher(u+"/save?table=userupdate", `{"id": "f89813e7-27b0-478d-b20e-f36a9ba9d759","email": "eb@rambler.ru"}`)
		}
	case 11:
		{
			Pusher(u+"/save?table=userdelete", `{"id": "f89813e7-27b0-478d-b20e-f36a9ba9d759"}`)
		}
	case 12:
		{
			Pusher(u+"/save?table=orderadd", `{"uid": "a30aaed1-7522-4c37-a6a9-b5774fe872e6"}`)
		}
	case 13:
		{
			Pusher(u+"/save?table=orderfill", `{"oid": 14, "pid":2}`)
		}
	case 14:
		{
			Pusher(u+"/save?table=orderdel", `{"oid": "15"}`)
		}
	}
}

func Getter(u string) {
	result, err := http.Get(u) //nolint
	if err != nil {
		fmt.Println(err)
		return
	}
	defer result.Body.Close()

	body, err := io.ReadAll(result.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func Pusher(u string, d string) {
	result, err := http.Post(u, "", bytes.NewReader([]byte(d))) //nolint
	if err != nil {
		fmt.Println(err)
		return
	}
	defer result.Body.Close()

	body, err := io.ReadAll(result.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
