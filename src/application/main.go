package main

import (
	"awesomeProject/src/handle"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", handle.GetPeople).Methods("GET")
	router.HandleFunc("/person/", handle.InsertPerson).Methods("POST")
	router.HandleFunc("/person/", handle.UpdatePerson).Methods("PUT")
	router.HandleFunc("/person/", handle.DeletePerson).Methods("DELETE")
	router.HandleFunc("/users/filter/", handle.UsersFilter).Methods("GET")
	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
