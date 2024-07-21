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
	http.HandleFunc("/get", handlerGet)
	server := &http.Server{
		Addr:              u,
		ReadHeaderTimeout: 3 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln("start failed")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Println(r.Method, r.RemoteAddr)
}

func handlerGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.RemoteAddr)
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no"))
		return
	}

	w.Write([]byte("This is a Get req"))
}

func handlerSave(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.RemoteAddr)
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
	w.Write(body)
}
