package Controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

var Database *sql.DB

type TProduct struct {
	Id      int
	Model   string
	Company string
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	rows, err := Database.Query("SELECT * FROM productdb.products")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	products := []TProduct{}
	for rows.Next() {
		p := TProduct{}
		rows.Scan(&p.Id, &p.Model, &p.Company)
		products = append(products, p)
	}
	json.NewEncoder(w).Encode(products)
}

func GetOneProductById(w http.ResponseWriter, r *http.Request) {
	paramsId := mux.Vars(r)
	rows, err := Database.Query("SELECT * FROM productdb.products WHERE id = ?", paramsId["id"])
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	products := []TProduct{}
	for rows.Next() {
		p := TProduct{}
		rows.Scan(&p.Id, &p.Model, &p.Company)
		products = append(products, p)
	}
	json.NewEncoder(w).Encode(products)
}

func CreateProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	smtp, err := Database.Prepare("INSERT INTO productdb.products (Model, Company) VALUES (?, ?)")
	if err != nil {
		fmt.Println(err)
	}
	body, err := io.ReadAll(r.Body)

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	Model := keyVal["Model"]
	Company := keyVal["Company"]
	_, err = smtp.Exec(Model, Company)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "Created post: Model = %s, Company = %s", Model, Company)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	paramsId := mux.Vars(r)
	smtp, err := Database.Prepare("DELETE FROM productdb.products WHERE id = ?")
	if err != nil {
		fmt.Println(err)
	}
	_, err = smtp.Exec(paramsId["id"])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "Delete post with id = %s", paramsId["id"])
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	paramsId := mux.Vars(r)

	smtp, err := Database.Prepare("UPDATE productdb.products SET Model = ?, Company = ? WHERE id = ?")
	if err != nil {
		log.Println(err)
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	Model := keyVal["Model"]
	Company := keyVal["Company"]

	_, err = smtp.Exec(Model, Company, paramsId["id"])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "Update Post ID = %s; Model = %s, Company = %s", paramsId["id"], Model, Company)
}
