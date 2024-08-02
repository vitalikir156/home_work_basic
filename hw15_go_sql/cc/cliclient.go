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
			result, err := http.Get(u + "/get") //nolint
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
	case 2:
		{
			result, err := http.Get(u + "/get?table=users") //nolint
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
	case 3:
		{
			result, err := http.Get(u + "/get?table=userorders&srch=a30aaed1-7522-4c37-a6a9-b5774fe872e6") //nolint
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
	case 4:
		{
			result, err := http.Get(u + "/get?table=usersumm&srch=a342e4da-1438-4d60-bd4e-7beef2791324") //nolint
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
	case 5:
		{
			result, err := http.Get(u + "/get?table=useraver&srch=a342e4da-1438-4d60-bd4e-7beef2791324") //nolint
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
	case 6:
		{
			data := `{"desc": "ATTINY 13, 8b, 9.6mhz 1k flash","price": 170}`

			result, err := http.Post(u+"/save", "", bytes.NewReader([]byte(data))) //nolint
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
	case 7:
		{
			data := `{ "id":4, "desc": "ATTINY 13, 8b, 9.6mhz 1k flash","price": 170}`

			result, err := http.Post(u+"/save?table=tableupdate", "", bytes.NewReader([]byte(data))) //nolint
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
	case 8:
		{
			data := `{ "id":9}`

			result, err := http.Post(u+"/save?table=deleterow", "", bytes.NewReader([]byte(data))) //nolint
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
	case 9:
		{
			data := `{"name": "Edward Bill","email": "eb@rambler.ru", "password":"1212"}`

			result, err := http.Post(u+"/save?table=usercreate", "", bytes.NewReader([]byte(data))) //nolint
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
	case 10:
		{
			data := `{"id": "f89813e7-27b0-478d-b20e-f36a9ba9d759","email": "eb@rambler.ru"}`

			result, err := http.Post(u+"/save?table=userupdate", "", bytes.NewReader([]byte(data))) //nolint
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
	case 11:
		{
			data := `{"id": "f89813e7-27b0-478d-b20e-f36a9ba9d759"}`

			result, err := http.Post(u+"/save?table=userdelete", "", bytes.NewReader([]byte(data))) //nolint
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
	case 12:
		{
			data := `{"uid": "a30aaed1-7522-4c37-a6a9-b5774fe872e6"}`

			result, err := http.Post(u+"/save?table=orderadd", "", bytes.NewReader([]byte(data))) //nolint
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
	case 13:
		{
			data := `{"oid": 14, "pid":2}`

			result, err := http.Post(u+"/save?table=orderfill", "", bytes.NewReader([]byte(data))) //nolint
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
	case 14:
		{
			data := `{"oid": "15"}`

			result, err := http.Post(u+"/save?table=orderdel", "", bytes.NewReader([]byte(data))) //nolint
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
	}
}
