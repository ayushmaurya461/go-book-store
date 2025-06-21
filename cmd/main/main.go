package main

import (
	"log"
	"net/http"

	"github.com/ayushmaurya461/go-book-store.git/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
