package sqltranslation

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgress support
)

func Start(config string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", config)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(6)
	db.SetMaxIdleConns(6)
	return db
}

func GetProducts(db *sqlx.DB) (string, error) {
	query := `select * from market.products`
	rows, err := db.QueryContext(context.Background(), query)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	var out strings.Builder
	for rows.Next() {
		var id int64
		var price float64
		var name string
		if err = rows.Scan(&id, &name, &price); err != nil {
			fmt.Println(err)
		} else {
			out.WriteString(fmt.Sprintf("%d %v %v\n", id, name, price))
		}
	}
	return out.String(), nil
}

func GetUsers(db *sqlx.DB) (string, error) {
	query := `select id, name, email from market.users`
	rows, err := db.QueryContext(context.Background(), query)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	var out strings.Builder
	for rows.Next() {
		var id string
		var email string
		var name string
		if err = rows.Scan(&id, &name, &email); err != nil {
			fmt.Println(err)
		} else {
			out.WriteString(fmt.Sprintf("%v %v %v\n", id, name, email))
		}
	}
	return out.String(), nil
}

func GetUserOrders(db *sqlx.DB, q string) (string, error) {
	query := `select user_id, users.name, order_date, orders.id, total_amount, price,
	products.name from ((market.orderproducts join market.orders on id=orderid) join
	market.products on products.id=productid) join market.users on users.id = user_id where user_id=$1`
	rows, err := db.QueryContext(context.Background(), query, q)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	var out strings.Builder
	for rows.Next() {
		var id, name, odate, pname string
		var tam, price float64
		var oid int64
		if err = rows.Scan(&id, &name, &odate, &oid, &tam, &price, &pname); err != nil {
			fmt.Println(err)
		} else {
			out.WriteString(fmt.Sprintf("%v %v %v %v %v %v %v\n", id, name, odate, oid, tam, price, pname))
		}
	}
	return out.String(), nil
}

func GetUserSumm(db *sqlx.DB, q string) (string, error) {
	query := `SELECT subqw.uid, subqw.uname, COALESCE(SUM(subqw.total), 0) AS sum_price
	FROM (select user_id as uid, users.name as uname, order_date, orders.id as "order id",
	 total_amount as total from market.orders  join market.users on users.id = user_id) as subqw where uid=$1
	GROUP BY subqw.uid, subqw.uname`
	rows, err := db.QueryContext(context.Background(), query, q)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	var out strings.Builder
	for rows.Next() {
		var id string
		var price float64
		var name string
		if err = rows.Scan(&id, &name, &price); err != nil {
			fmt.Println(err)
		} else {
			out.WriteString(fmt.Sprintf("%v %v %v\n", id, name, price))
		}
	}
	return out.String(), nil
}

func GetUserAver(db *sqlx.DB, q string) (string, error) {
	query := `SELECT subqw.uid, subqw.uname, COALESCE(AVG(subqw.prodprice), 0) AS avg_price
	FROM (select user_id as uid, users.name as uname, order_date, orders.id, total_amount, price as prodprice,
	products.name from ((market.orderproducts join market.orders on id=orderid) 
	join market.products on products.id=productid)
	join market.users on users.id = user_id) as subqw where uid=$1
	GROUP BY subqw.uid, subqw.uname`
	rows, err := db.QueryContext(context.Background(), query, q)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	var out strings.Builder
	for rows.Next() {
		var id string
		var price float64
		var name string
		if err = rows.Scan(&id, &name, &price); err != nil {
			fmt.Println(err)
		} else {
			out.WriteString(fmt.Sprintf("%v %v %v\n", id, name, price))
		}
	}
	return out.String(), nil
}

func PostNewProduct(db *sqlx.DB, q []byte) (string, error) {
	var res struct {
		Desc  string  `json:"desc"`
		Price float64 `json:"price"`
	}
	err := json.Unmarshal(q, &res)
	if err != nil {
		return "", err
	}
	tx, err := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	defer tx.Rollback()
	if err != nil {
		return fmt.Sprintf("%v", err), err
	}
	out, err := tx.Exec("INSERT INTO market.products (name, price) VALUES ($1, $2)", res.Desc, res.Price)
	if err != nil {
		return fmt.Sprintf("%v", err), err
	}
	out2, err := out.RowsAffected()
	if err != nil {
		return fmt.Sprintf("%v", err), err
	}
	err = tx.Commit()
	return fmt.Sprintf("%v line(s) inserted", out2), err
}

func PostEditProduct(db *sqlx.DB, q []byte) (string, error) {
	var res struct {
		ID    int64   `json:"id"`
		Desc  string  `json:"desc"`
		Price float64 `json:"price"`
	}
	err := json.Unmarshal(q, &res)
	if err != nil {
		return "", err
	}
	tx, err := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	defer tx.Rollback()
	if err != nil {
		return fmt.Sprintf("%v", err), err
	}
	query := `SELECT COUNT(*) AS total_rows	FROM market.products where id=$1`
	check := tx.QueryRowContext(context.Background(), query, res.ID)
	var row int64
	err = check.Scan(&row)
	if err != nil {
		return fmt.Sprintf("%v", err), err
	}
	if row < 1 {
		return "selected ID not found", fmt.Errorf("id not found")
	}
	if len(res.Desc) > 0 {
		_, err = tx.Exec("update market.products set name=$1 where id=$2", res.Desc, res.ID)
		if err != nil {
			return fmt.Sprintf("%v", err), err
		}
	}
	if res.Price > 0 {
		_, err = tx.Exec("update market.products set price=$1 where id=$2", res.Price, res.ID)
		if err != nil {
			return fmt.Sprintf("%v", err), err
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Sprintf("%v", err), err
	}
	return "update success", err
}

func PostDelProduct(db *sqlx.DB, q []byte) (string, error) {
	var res struct {
		ID int64 `json:"id"`
	}
	err := json.Unmarshal(q, &res)
	if err != nil {
		return "", err
	}
	tx, err := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	defer tx.Rollback()
	if err != nil {
		return fmt.Sprintf("%v", err), err
	}
	out, err := tx.Exec("delete from market.products where id=$1", res.ID)
	if err != nil {
		return fmt.Sprintf("%v", err), err
	}
	out2, err := out.RowsAffected()
	if err != nil {
		return fmt.Sprintf("%v", err), err
	}
	err = tx.Commit()
	return fmt.Sprintf("%v line(s) deleted", out2), err
}

func PostNewUser(db *sqlx.DB, q []byte) (string, error) {
	var res struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := json.Unmarshal(q, &res)
	if err != nil {
		return "", err
	}
	tx, err := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	defer tx.Rollback()
	if err != nil {
		return fmt.Sprintf("%v", err), err
	}
	query := `INSERT INTO market.users (name, email,password) VALUES ($1, $2, crypt($3, gen_salt('md5'))) returning id`
	out := tx.QueryRow(query, res.Name, res.Email, res.Password)
	var row string
	out.Scan(&row)
	err = tx.Commit()
	if err != nil {
		return fmt.Sprintf("%v", err), err
	}
	return fmt.Sprintf("uuid= %v", row), err
}

func PostEditUser(db *sqlx.DB, q []byte) (string, error) {
	var res struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := json.Unmarshal(q, &res)
	if err != nil {
		return "", err
	}
	tx, err := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	defer tx.Rollback()
	if err != nil {
		return fmt.Sprintf("%v", err), err
	}
	query := `SELECT COUNT(*) AS total_rows	FROM market.users where id=$1`
	check := tx.QueryRowContext(context.Background(), query, res.ID)
	var row int64
	err = check.Scan(&row)
	if err != nil {
		return fmt.Sprintf("%v", err), err
	}
	if row < 1 {
		return "selected ID not found", fmt.Errorf("id not found")
	}
	if len(res.Name) > 0 {
		_, err = tx.Exec("update market.users set name=$1 where id=$2", res.Name, res.ID)
		if err != nil {
			return fmt.Sprintf("%v", err), err
		}
	}
	if len(res.Email) > 0 {
		_, err = tx.Exec("update market.users set email=$1 where id=$2", res.Email, res.ID)
		if err != nil {
			return fmt.Sprintf("%v", err), err
		}
	}
	if len(res.Password) > 0 {
		_, err = tx.Exec("update market.users set password=crypt($3, gen_salt('md5')) where id=$2", res.Password, res.ID)
		if err != nil {
			return fmt.Sprintf("%v", err), err
		}
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Sprintf("%v", err), err
	}
	return "update with no errors", err
}

func PostDelUser(db *sqlx.DB, q []byte) (string, error) {
	var res struct {
		ID string `json:"id"`
	}
	err := json.Unmarshal(q, &res)
	if err != nil {
		return "", err
	}
	tx, err := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	defer tx.Rollback()
	if err != nil {
		return fmt.Sprintf("%v", err), err
	}
	out, err := tx.Exec("delete from market.users where id=$1", res.ID)
	if err != nil {
		return fmt.Sprintf("%v", err), err
	}
	out2, err := out.RowsAffected()
	if err != nil {
		return fmt.Sprintf("%v", err), err
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Sprintf("%v", err), err
	}
	return fmt.Sprintf("%v line(s) deleted", out2), err
}

func PostCreateOrder(db *sqlx.DB, q []byte) (string, error) {
	var res struct {
		UID string `json:"uid"`
	}
	err := json.Unmarshal(q, &res)
	if err != nil {
		return "", err
	}
	tx, err := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	defer tx.Rollback()
	if err != nil {
		return fmt.Sprintf("%v", err), err
	}
	query1 := `SELECT COUNT(*) AS total_rows	FROM market.users where id=$1`
	check := tx.QueryRowContext(context.Background(), query1, res.UID)
	var row2 int64
	err = check.Scan(&row2)
	if err != nil {
		return fmt.Sprintf("%v", err), err
	}
	if row2 < 1 {
		return "selected user ID not found", fmt.Errorf("user id not found")
	}
	query := `INSERT INTO market.orders (user_id, order_date, total_amount) VALUES ($1, current_date, 0) returning id`
	out := tx.QueryRow(query, res.UID)
	var row string
	out.Scan(&row)
	err = tx.Commit()
	if err != nil {
		return fmt.Sprintf("%v", err), err
	}
	return fmt.Sprintf("uuid= %v", row), err
}

func PostDelOrder(db *sqlx.DB, q []byte) (string, error) {
	var res struct {
		ID string `json:"oid"`
	}
	err := json.Unmarshal(q, &res)
	if err != nil {
		return "", err
	}
	fmt.Printf("")
	tx, err := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	defer tx.Rollback()
	if err != nil {
		return fmt.Sprintf("%v", err), err
	}
	out, err := tx.Exec("delete from market.orders where id=$1", res.ID)
	if err != nil {
		return fmt.Sprintf("%v", err), err
	}
	out2, err := out.RowsAffected()
	if err != nil {
		return fmt.Sprintf("%v", err), err
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Sprintf("%v", err), err
	}
	return fmt.Sprintln(" line(s) deleted", out2), err
}

func PostAddOrdProd(db *sqlx.DB, q []byte) (string, error) {
	var res struct {
		Oid int64 `json:"oid"`
		Pid int64 `json:"pid"`
	}
	err := json.Unmarshal(q, &res)
	if err != nil {
		return "", err
	}
	tx, err := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	defer tx.Rollback()
	if err != nil {
		return fmt.Sprintf("%v", err), err
	}
	out, err := tx.Exec("INSERT INTO market.OrderProducts (orderId, ProductId) VALUES ($1, $2)", res.Oid, res.Pid)
	if err != nil {
		return fmt.Sprintf("%v", err), err
	}
	out2, err := out.RowsAffected()
	if err != nil {
		return fmt.Sprintf("%v", err), err
	}
	query2 := `update market.orders 
	SET total_amount = subqw.avg_price from (SELECT market.OrderProducts.orderId as oid,
	COALESCE(SUM(market.products.price), 0) AS avg_price FROM market.OrderProducts LEFT JOIN market.products 
	ON products.id=productid GROUP BY oid) as subqw
	where market.orders.id = subqw.oid and market.orders.id=$1`
	_, err = tx.Exec(query2, res.Oid)
	if err != nil {
		return fmt.Sprintf("%v", err), err
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Sprintf("%v", err), err
	}
	return fmt.Sprintf("%v line(s) inserted", out2), err
}
