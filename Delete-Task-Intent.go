package main

import (
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
)

//DeleteTaskIntent ...
type DeleteTaskIntent struct {
	TaskRepository TaskRepository
}

//Enact ...
func (i DeleteTaskIntent) Enact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	val, err := i.TaskRepository.Delete(vars["id"])
	if err == nil && val != 0 {
		fmt.Fprintf(w, "delete task succesfully")
		fmt.Println(err, val)
	} else {
		fmt.Fprintf(w, "delete task unsuccesfully")
		fmt.Println(err, val)
	}
}
