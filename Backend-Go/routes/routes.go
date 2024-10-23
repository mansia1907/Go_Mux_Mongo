package routes

import (
	"Backend-Go/userhandler"

	"github.com/gorilla/mux"
)

func Routes(router *mux.Router) {
	//All routes related to users comes here

	router.HandleFunc("/api/signup", userhandler.SignUpHandler()).Methods("POST") //add this

}
