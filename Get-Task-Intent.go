package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetTaskIntent used to retrieve a single task from the repository.
type GetTaskIntent struct {
	ListRepository            ListRepository
	validateTaskSpecification ValidateTaskSpecification
}

// Enact method takes task id as URL parameter and serves a task object.
func (i GetTaskIntent) Enact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := i.validateTaskSpecification.Enact(vars["taskid"], vars["listid"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mocklist, err := i.ListRepository.GetTaskFromList(vars["taskid"])
	if mocklist.ID == 0 || err != nil {
		if err == nil {
			http.Error(w, "Error occured while retreival", http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		data, err := json.Marshal(mocklist)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		return
	}
}
