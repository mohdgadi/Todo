package main

import (
	"encoding/json"
	"fmt"

	"net/http"
)

// AddTaskIntent used to add task in the repository.
type AddTaskIntent struct {
	TaskRepository TaskRepository
	ListRepository ListRepository
}

// Enact takes JSON request and add task to the list.
func (i AddTaskIntent) Enact(w http.ResponseWriter, r *http.Request) {
	var (
		reqJSON List
		task    Task
	)
	if err := json.NewDecoder(r.Body).Decode(&reqJSON); err == nil {
		task = reqJSON.Tasks[0]
		check := i.ListRepository.Check(reqJSON.Name)
		if check == false {
			http.Error(w, "List doesnt exist", http.StatusBadRequest)
			return
		}
		err = i.TaskRepository.Add(task, reqJSON.Name)
		if err == nil {
			fmt.Fprintf(w, "Task added succesfully")
			w.WriteHeader(http.StatusOK)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
