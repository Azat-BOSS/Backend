package main

import (
	"Natural/Database"
	"Natural/Routers"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func StartServer() {
	fmt.Println("Server is working")
	http.ListenAndServe("localhost:8080", nil)
}

func main() {
	Database.StartMySQL()
	Routers.ProductRouter()
	StartServer()
}
