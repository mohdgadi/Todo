package main

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
)

// UpdateTaskIntent used to update task status in the database.
type UpdateTaskIntent struct {
	ListRepository            ListRepository
	ValidateTaskSpecification ValidateTaskSpecification
}

// Enact method is usedt change the status of task completed .
func (i UpdateTaskIntent) Enact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := i.ValidateTaskSpecification.Enact(vars["taskid"], vars["listid"])
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	var reqJSON Task
	if err = json.NewDecoder(r.Body).Decode(&reqJSON); err == nil {
		status := reqJSON.Status
		err := i.ListRepository.UpdateTask(vars["listid"], vars["taskid"], status)
		if err == nil {
			fmt.Fprintf(w, "Task updated succesfully")
			w.WriteHeader(http.StatusOK)
		} else {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}
	} else {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
}
