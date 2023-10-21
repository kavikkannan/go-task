package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kavikkannan/go-task/pkg/routes"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStors(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}