package main

import (
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
)

// DeleteTaskIntent used to delete a single task from the reposiotry.
type DeleteTaskIntent struct {
	ListRepository            ListRepository
	ValidateTaskSpecification ValidateTaskSpecification
}

// Enact takes task ID as URL parameter deletes a singular task by ID.
func (i DeleteTaskIntent) Enact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := i.ValidateTaskSpecification.Enact(vars["taskid"], vars["listid"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = i.ListRepository.DeleteTaskFromList(vars["taskid"])
	if err == nil {
		fmt.Fprintf(w, "delete task succesfully")
		return
	}
	http.Error(w, err.Error(), http.StatusBadRequest)
}
