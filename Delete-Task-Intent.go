package main

import (
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
)

// DeleteTaskIntent used to delete a single task from the reposiotry.
type DeleteTaskIntent struct {
	TaskRepository TaskRepository
}

// Enact takes task ID as URL parameter deletes a singular task by ID.
func (i DeleteTaskIntent) Enact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := i.TaskRepository.Delete(vars["id"])
	if err == nil {
		fmt.Fprintf(w, "delete task succesfully")
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
