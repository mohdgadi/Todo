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
	err := i.TaskRepository.Delete(vars["id"])
	if err == nil {
		fmt.Fprintf(w, "delete task succesfully")

	} else {
		fmt.Fprintf(w, err.Error())

	}
}
