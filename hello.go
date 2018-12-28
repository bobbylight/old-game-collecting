package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/api/people", GetPeople).Methods("GET")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	log.Println("Listening...")
	http.ListenAndServe(":3000", router)
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world again")
}
