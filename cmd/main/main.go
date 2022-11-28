package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/J-hon/bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)

	fmt.Printf("Starting server at port 9010\n")

	log.Fatal(http.ListenAndServe(":9010", r))
}
