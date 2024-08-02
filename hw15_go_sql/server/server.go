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

func Server(u string) {
	db = sqltranslation.Start("user=vit password=password2717 dbname=market port=5432 sslmode=disable")
	defer db.Close()
	http.HandleFunc("/", handler)
	http.HandleFunc("/save", handlerSave)
	http.HandleFunc("/get", handlerGet)
	server := &http.Server{
		Addr:              u,
		ReadHeaderTimeout: 3 * time.Second,
	}
	fmt.Println("ready to start")
	if err := server.ListenAndServe(); err != nil {
		return
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Println(r.Method, r.RemoteAddr)
}

func handlerGet(w http.ResponseWriter, r *http.Request) {
	tname := r.URL.Query().Get("table")
	srch := r.URL.Query().Get("srch")
	fmt.Println(r.Method, r.RemoteAddr)
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no"))
		return
	}
	switch tname {
	case "users":
		{
			out, err := sqltranslation.GetUsers(db)
			if err != nil {
				w.Write([]byte(fmt.Sprintf("%v", err)))
			}

			w.Write([]byte(out))
		}
	case "userorders":
		{
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
	case "usersumm":
		{
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
	case "useraver":
		{
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
	default:
		{
			out, err := sqltranslation.GetProducts(db)
			if err != nil {
				w.Write([]byte(fmt.Sprintf("%v", err)))
			}
			w.Write([]byte(out))
		}
	}
}

func handlerSave(w http.ResponseWriter, r *http.Request) {
	tname := r.URL.Query().Get("table")
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
	switch tname {
	case "tableupdate":
		{
			out, err2 := sqltranslation.PostEditProduct(db, body)
			if err2 != nil {
				fmt.Println(err2)
				w.Write([]byte(err2.Error()))
				return
			}
			w.Write([]byte(out))
		}
	case "deleterow":
		{
			out, err2 := sqltranslation.PostDelProduct(db, body)
			if err2 != nil {
				fmt.Println(err2)
				w.Write([]byte(err2.Error()))
				return
			}
			w.Write([]byte(out))
		}
	case "userupdate": //-
		{
			out, err2 := sqltranslation.PostEditUser(db, body)
			if err2 != nil {
				fmt.Println(err2)
				w.Write([]byte(err2.Error()))
				return
			}
			w.Write([]byte(out))
		}
	case "userdelete": //-
		{
			out, err2 := sqltranslation.PostDelUser(db, body)
			if err2 != nil {
				fmt.Println(err2)
				w.Write([]byte(err2.Error()))
				return
			}
			w.Write([]byte(out))
		}
	case "usercreate":
		{
			out, err2 := sqltranslation.PostNewUser(db, body)

			if err2 != nil {
				fmt.Println(err2)
				w.Write([]byte(err2.Error()))
				return
			}
			w.Write([]byte(out))
		}
	case "tableadd":
		{
			out, err2 := sqltranslation.PostNewProduct(db, body)

			if err2 != nil {
				fmt.Println(err2)
				w.Write([]byte(err2.Error()))
				return
			}
			w.Write([]byte(out))
		}
	case "orderadd":
		{
			out, err2 := sqltranslation.PostCreateOrder(db, body)

			if err2 != nil {
				fmt.Println(err2)
				w.Write([]byte(err2.Error()))
				return
			}
			w.Write([]byte(out))
		}
	case "orderdel":
		{
			out, err2 := sqltranslation.PostDelOrder(db, body)

			if err2 != nil {
				fmt.Println(err2)
				w.Write([]byte(err2.Error()))
				return
			}
			w.Write([]byte(out))
		}
	case "orderfill":
		{
			out, err2 := sqltranslation.PostAddOrdProd(db, body)

			if err2 != nil {
				fmt.Println(err2)
				w.Write([]byte(err2.Error()))
				return
			}
			w.Write([]byte(out))
		}
	default:
		{
			w.Write([]byte("wrong request"))
		}
	}
}
