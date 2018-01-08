package main

import (
	"encoding/json"
	"strconv"

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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var task Task
	if err = json.NewDecoder(r.Body).Decode(&task); err == nil {
		task.ID, err = strconv.Atoi(vars["taskid"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		list := List{Tasks: []Task{task}}
		err := i.ListRepository.UpdateTaskStatus(list)
		if err == nil {
			w.WriteHeader(http.StatusOK)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
