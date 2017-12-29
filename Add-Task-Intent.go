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
		reqJSON List
	)
	if err := json.NewDecoder(r.Body).Decode(&reqJSON); err == nil {
		reqJSON.Name = vars["listid"]
		err = i.ListRepository.AddTaskToList(reqJSON)
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
