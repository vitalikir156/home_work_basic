package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func Server(u string) {
	http.HandleFunc("/", handler)
	http.HandleFunc("/save", handlerSave)

	server := &http.Server{
		Addr:              u,
		ReadHeaderTimeout: 3 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln("start failed")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello there!")
	fmt.Println(r.Method, r.RemoteAddr)
}

func handlerSave(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no"))
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r.Method, r.RemoteAddr)
	w.Write(body)
}
