package server

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	sqltranslation "github.com/vitalikir156/home_work_basic/hw15_go_sql/dbc"
)

var db *sqlx.DB

func HTTP(u string, dbin *sqlx.DB) {
	db = dbin
	http.HandleFunc("/", handler)
	http.HandleFunc("GET /products", handlerGetProd)
	http.HandleFunc("GET /users", handlerGetUsers)
	http.HandleFunc("GET /user/aver", handlerGetUserAver)
	http.HandleFunc("GET /user/summ", handlerGetUserSumm)
	http.HandleFunc("GET /orders", handlerGetOrders)
	http.HandleFunc("POST /products", handlerPostProd)
	http.HandleFunc("POST /users", handlerPostUsers)
	http.HandleFunc("POST /orders", handlerPostOrders)
	http.HandleFunc("POST /orderproducts", handlerPostOrderProducts)
	http.HandleFunc("UPDATE /products", handlerUpdProd)
	http.HandleFunc("UPDATE /users", handlerUpdUser)
	http.HandleFunc("DELETE /products", handlerDelProducts)
	http.HandleFunc("DELETE /users", handlerDelUser)
	http.HandleFunc("DELETE /orders", handlerDelOrder)

	server := &http.Server{
		Addr:              u,
		ReadHeaderTimeout: 3 * time.Second,
	}
	fmt.Println("ready to start")
	if err := server.ListenAndServe(); err != nil {
		return
	}
}

func handlerGetProd(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.RemoteAddr)
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no"))
		return
	}
	out, err := sqltranslation.GetProducts(db)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	w.Write([]byte(out))
}

func handlerGetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.RemoteAddr)
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no"))
		return
	}
	out, err := sqltranslation.GetUsers(db)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}

	w.Write([]byte(out))
}

func handlerGetUserAver(w http.ResponseWriter, r *http.Request) {
	srch := r.URL.Query().Get("srch")
	fmt.Println(r.Method, r.RemoteAddr)
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no"))
		return
	}
	if len(srch) == 0 {
		w.Write([]byte("No search request provided!"))
	} else {
		out, err := sqltranslation.GetUserAver(db, srch)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("%v", err)))
		}
		w.Write([]byte(out))
	}
}

func handlerGetUserSumm(w http.ResponseWriter, r *http.Request) {
	srch := r.URL.Query().Get("srch")
	fmt.Println(r.Method, r.RemoteAddr)
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no"))
		return
	}
	if len(srch) == 0 {
		w.Write([]byte("No search request provided!"))
	} else {
		out, err := sqltranslation.GetUserSumm(db, srch)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("%v", err)))
		}
		w.Write([]byte(out))
	}
}

func handlerGetOrders(w http.ResponseWriter, r *http.Request) {
	srch := r.URL.Query().Get("srch")
	fmt.Println(r.Method, r.RemoteAddr)
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no"))
		return
	}
	if len(srch) == 0 {
		w.Write([]byte("No search request provided!"))
	} else {
		out, err := sqltranslation.GetUserOrders(db, srch)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("%v", err)))
		}
		w.Write([]byte(out))
	}
}

func handlerPostProd(w http.ResponseWriter, r *http.Request) {
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

	out, err2 := sqltranslation.PostNewProduct(db, body)

	if err2 != nil {
		fmt.Println(err2)
		w.Write([]byte(err2.Error()))
		return
	}
	w.Write([]byte(out))
}

func handlerPostUsers(w http.ResponseWriter, r *http.Request) {
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
	out, err2 := sqltranslation.PostNewUser(db, body)

	if err2 != nil {
		fmt.Println(err2)
		w.Write([]byte(err2.Error()))
		return
	}
	w.Write([]byte(out))
}

func handlerPostOrders(w http.ResponseWriter, r *http.Request) {
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
	out, err2 := sqltranslation.PostCreateOrder(db, body)
	if err2 != nil {
		fmt.Println(err2)
		w.Write([]byte(err2.Error()))
		return
	}
	w.Write([]byte(out))
}

func handlerPostOrderProducts(w http.ResponseWriter, r *http.Request) {
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
	out, err2 := sqltranslation.PostAddOrdProd(db, body)
	if err2 != nil {
		fmt.Println(err2)
		w.Write([]byte(err2.Error()))
		return
	}
	w.Write([]byte(out))
}

func handlerUpdProd(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.RemoteAddr)
	if r.Method != "UPDATE" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no"))
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	out, err2 := sqltranslation.PostEditProduct(db, body)
	if err2 != nil {
		fmt.Println(err2)
		w.Write([]byte(err2.Error()))
		return
	}
	w.Write([]byte(out))
}

func handlerUpdUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.RemoteAddr)
	if r.Method != "UPDATE" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no"))
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	out, err2 := sqltranslation.PostEditUser(db, body)
	if err2 != nil {
		fmt.Println(err2)
		w.Write([]byte(err2.Error()))
		return
	}
	w.Write([]byte(out))
}

func handlerDelProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.RemoteAddr)
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no"))
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	out, err2 := sqltranslation.PostDelProduct(db, body)
	if err2 != nil {
		fmt.Println(err2)
		w.Write([]byte(err2.Error()))
		return
	}
	w.Write([]byte(out))
}

func handlerDelUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.RemoteAddr)
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no"))
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	out, err2 := sqltranslation.PostDelUser(db, body)
	if err2 != nil {
		fmt.Println(err2)
		w.Write([]byte(err2.Error()))
		return
	}
	w.Write([]byte(out))
}

func handlerDelOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.RemoteAddr)
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no"))
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	out, err2 := sqltranslation.PostDelOrder(db, body)

	if err2 != nil {
		fmt.Println(err2)
		w.Write([]byte(err2.Error()))
		return
	}
	w.Write([]byte(out))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Println(r.Method, r.RemoteAddr)
}
