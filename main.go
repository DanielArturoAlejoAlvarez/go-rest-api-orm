package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World!</h1>")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users/{name}", getUser).Methods("GET")
	router.HandleFunc("/users/{name}/{email}", saveUser).Methods("POST")
	router.HandleFunc("/users/{name}/{email}", updateUser).Methods("PUT")
	router.HandleFunc("/users/{name}", deleteUser).Methods("DELETE")
	router.HandleFunc("/", helloWorld).Methods("GET")
	log.Fatal(http.ListenAndServe(":5500", router))
}

func main() {
	fmt.Println("Welcome to REST API with ORM")

	InitialMigration()
	handleRequests()
}
