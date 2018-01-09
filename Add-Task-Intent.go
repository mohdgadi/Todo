package main

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
)

// AddTaskIntent used to add task in the repository.
type AddTaskIntent struct {
	ListRepository ListRepository
}

// Enact takes JSON request and add task to the list.
func (i AddTaskIntent) Enact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var (
		task Task
	)
	if err := json.NewDecoder(r.Body).Decode(&task); err == nil {
		if task.Name == "" {
			http.Error(w, "Task cannot be empty", http.StatusBadRequest)
			return
		}
		list := List{Tasks: []Task{task}, Name: vars["listid"]}
		if i.ListRepository.CheckIfExists(list.Name) == false {
			http.Error(w, "List doesnt exists", http.StatusBadRequest)
			return
		}
		err = i.ListRepository.AddTaskToList(list)
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
