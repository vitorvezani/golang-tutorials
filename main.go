package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/todos", handleTodos)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleTodos(w http.ResponseWriter, r *http.Request) {
	person := Person{
		Id:   1,
		Name: "Vitor",
	}

	changePersonData(&person)

	log.Printf("person is: %v", person)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(person); err != nil {
		log.Default().Printf("could not write json: %v", err)
	}
}

func changePersonData(person *Person) {
	person.Id = 2
	person.Name = "José"
}
