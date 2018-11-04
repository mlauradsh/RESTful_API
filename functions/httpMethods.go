package functions

import (
	"awesomeProject/constructors"
	// const "awesomeProject/constructors" to rename the import package path
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// Display all from people
func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

// Display a single selected person from people
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, person := range people {
		if person.ID == params["id"] {
			// Display the found person
			json.NewEncoder(w).Encode(person)
			return
		}

	}
	// If it doesn't find it in the array, then show No Content message
	w.WriteHeader(204)
	//w.Write([]byte("The person " + params["id"] + " has not been found in people"))
}

// Delete a person from people
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, person := range people {
		if person.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("The person " + person.ID + " has been deleted"))
			return
		}
	}
	w.WriteHeader(http.StatusNoContent)
	//w.Write([]byte("The person has not been found in people"))
	//json.NewEncoder(w).Encode(people)
}

// Create a person
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person constructors.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(person)
}
