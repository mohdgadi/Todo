package main

import (
	"encoding/json"
	"fmt"

	"net/http"
)

// UpdateTaskIntent used to update task status in the database.
type UpdateTaskIntent struct {
	TaskRepository TaskRepository
}

// Enact method is usedt change the status of task completed .
func (i UpdateTaskIntent) Enact(w http.ResponseWriter, r *http.Request) {
	var reqJSON Task
	if err := json.NewDecoder(r.Body).Decode(&reqJSON); err == nil {
		err := i.TaskRepository.Update(reqJSON)
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
