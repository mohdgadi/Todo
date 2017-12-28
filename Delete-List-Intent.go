package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// DeleteListIntent used to delete list of tasks from the repository.
type DeleteListIntent struct {
	ListRepository ListRepository
}

// Enact gets listname from URL parameter and deletes the list.
func (i DeleteListIntent) Enact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := i.ListRepository.Delete(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		fmt.Fprintf(w, "delete list succesfully")
		w.WriteHeader(http.StatusOK)
	}
}
