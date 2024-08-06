package main

import (
	"fmt"
	"sync"

	"github.com/spf13/pflag"
	"github.com/vitalikir156/home_work_basic/hw15_go_sql/client" // бесполезный клиент
	sqltranslation "github.com/vitalikir156/home_work_basic/hw15_go_sql/dbc"
	"github.com/vitalikir156/home_work_basic/hw15_go_sql/server"
)

var wg sync.WaitGroup

func srv(url string) {
	db := sqltranslation.Start("user=vit password=password2717 dbname=market port=5432 sslmode=disable")
	defer db.Close()
	server.HTTP(url, db)
	wg.Done()
}

func main() {
	urlserv := pflag.StringP("Server URL", "s", "", "server url ':8989' for example")
	url := pflag.StringP("client URL", "c", "", "client url")
	mode := pflag.StringP("prog mode", "m", "", "s for server, c for client, d for dual")
	pflag.Parse()
	wg = sync.WaitGroup{}

	if *mode == "s" || *mode == "d" {
		if len(*urlserv) < 1 {
			fmt.Println("empty server url")
			return
		}
		wg.Add(1)
		go srv(*urlserv)
	}

	if *mode == "c" || *mode == "d" {
		if len(*url) < 1 {
			fmt.Println("empty client url")
			return
		}
		out, err := client.Client(*url)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(out)
	}

	wg.Wait()
}
