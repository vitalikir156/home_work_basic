package main

// типа клиент, потыкать серверную часть запросами.
import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
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
	1 Getter((u + "/products")) 
	2 Getter((u + "/users"))
	3 Getter((u + "/orders?srch=a30aaed1-7522-4c37-a6a9-b5774fe872e6"))
	4 Getter((u + "/user/summ?srch=a342e4da-1438-4d60-bd4e-7beef2791324"))
	5 Getter((u + "/user/aver?srch=a342e4da-1438-4d60-bd4e-7beef2791324"))
	6 Pusher("POST", u+"/products", {"desc": "ATTINY 13, 8b, 9.6mhz 1k flash","price": 170})
	7 Pusher("UPDATE", u+"/products", { "id":4, "desc": "ATTINY 13, 8b, 9.6mhz 1k flash","price": 170})
	8 Pusher("DELETE", u+"/products", { "id":9})
	9 Pusher("POST", u+"/users", {"name": "Edward Bill","email": "eb@rambler.ru", "password":"1212"})
	10 Pusher("UPDATE", u+"/users", {"id": "f89813e7-27b0-478d-b20e-f36a9ba9d759","email": "eb@rambler.ru"})
	11 Pusher("DELETE", u+"/users", {"id": "f89813e7-27b0-478d-b20e-f36a9ba9d759"})
	12 Pusher("POST", u+"/orders", {"uid": "a30aaed1-7522-4c37-a6a9-b5774fe872e6"})
	13 Pusher("POST", u+"/orderproducts", {"oid": 14, "pid":2})
	14 Pusher("DELETE", u+"/orders", {"oid": "15"})
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
			Getter((u + "/products"))
		}
	case 2:
		{
			Getter((u + "/users"))
		}
	case 3:
		{
			Getter((u + "/orders?srch=a30aaed1-7522-4c37-a6a9-b5774fe872e6"))
		}
	case 4:
		{
			Getter((u + "/user/summ?srch=a342e4da-1438-4d60-bd4e-7beef2791324"))
		}
	case 5:
		{
			Getter((u + "/user/aver?srch=a342e4da-1438-4d60-bd4e-7beef2791324"))
		}
	case 6:
		{
			Pusher("POST", u+"/products", `{"desc": "ATTINY 13, 8b, 9.6mhz 1k flash","price": 170}`)
		}
	case 7:
		{
			Pusher("UPDATE", u+"/products", `{ "id":4, "desc": "ATTINY 13, 8b, 9.6mhz 1k flash","price": 170}`)
		}
	case 8:
		{
			Pusher("DELETE", u+"/products", `{ "id":11}`)
		}
	case 9:
		{
			Pusher("POST", u+"/users", `{"name": "Edward Bill","email": "eb@rambler.ru", "password":"1212"}`)
		}
	case 10:
		{
			Pusher("UPDATE", u+"/users", `{"id": "f89813e7-27b0-478d-b20e-f36a9ba9d759","email": "eb@rambler.ru"}`)
		}
	case 11:
		{
			Pusher("DELETE", u+"/users", `{"id": "f89813e7-27b0-478d-b20e-f36a9ba9d759"}`)
		}
	case 12:
		{
			Pusher("POST", u+"/orders", `{"uid": "a30aaed1-7522-4c37-a6a9-b5774fe872e6"}`)
		}
	case 13:
		{
			Pusher("POST", u+"/orderproducts", `{"oid": 14, "pid":2}`)
		}
	case 14:
		{
			Pusher("DELETE", u+"/orders", `{"oid": "15"}`)
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

func Pusher(m, u, b string) {
	req, err := http.NewRequest(m, u, bytes.NewReader([]byte(b)))
	if err != nil {
		fmt.Println(err)
		return
	}
	client := http.Client{
		Timeout: 30 * time.Second,
	 }
	 result, err := client.Do(req)
	 if err != nil {
		fmt.Println(err)
		return
	}
	body, err := io.ReadAll(result.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
