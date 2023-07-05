package main

import(
	
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/aron-helu/Book_Management-Go/pkg/routes"
)

func main() {
	r:= httprouter.New()

	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))

}