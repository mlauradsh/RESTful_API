package main

import (
	//"crypto/dsa"
	//"encoding/json"
	//"fmt"
	"github.com/gorilla/mux"
	//"encoding/json"
	"awesomeProject/functions"
	"log"
	"net/http"
)

// our main function
func main() {

	router := mux.NewRouter()

	// add people
	functions.InitializePeople()

	// add endpoints
	router.HandleFunc("/people", functions.GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", functions.GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", functions.DeletePerson).Methods("DELETE")
	router.HandleFunc("/people/{id}", functions.CreatePerson).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
