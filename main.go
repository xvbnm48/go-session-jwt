package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xvbnm48/go-session-jwt/controllers/authcontroller"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
