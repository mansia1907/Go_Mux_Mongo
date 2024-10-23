package main

import (
	"log"
	"net/http"

	"Backend-Go/dbhandler"

	"Backend-Go/routes"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	dbhandler.Database_handler()

	routes.Routes(router) //add this

	log.Fatal(http.ListenAndServe(":6000", router))

}
