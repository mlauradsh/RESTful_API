package functions

import "awesomeProject/constructors"

var people []constructors.Person

func InitializePeople() {

	people = append(people, constructors.Person{ID: "1", Firstname: "Laura", Surname: "Dsh", Address: &constructors.Address{City: "London", State: "Greater London"}})
	people = append(people, constructors.Person{ID: "2", Firstname: "Dani", Surname: "Thp", Address: &constructors.Address{City: "London", State: "Greater London"}})
	people = append(people, constructors.Person{ID: "3", Firstname: "Roberto", Surname: "Garcia", Address: &constructors.Address{City: "London", State: "Greater London"}})
	people = append(people, constructors.Person{ID: "4", Firstname: "Sarah", Surname: "Connor", Address: &constructors.Address{City: "London", State: "Greater London"}})
	people = append(people, constructors.Person{ID: "5", Firstname: "Lara", Surname: "Blue", Address: &constructors.Address{City: "London", State: "Greater London"}})
	people = append(people, constructors.Person{ID: "6", Firstname: "Norah", Surname: "Jones", Address: &constructors.Address{City: "London", State: "Greater London"}})
	people = append(people, constructors.Person{ID: "7", Firstname: "Alex", Surname: "Thomas", Address: &constructors.Address{City: "London", State: "Greater London"}})
	people = append(people, constructors.Person{ID: "8", Firstname: "Marcos", Surname: "Lopez", Address: &constructors.Address{City: "London", State: "Greater London"}})
}

//
//func GetPeopleData() []constructors.Person {
//	if (len(people) == 0) {
//		InitializePeople()
//	}
//	return people
//}
