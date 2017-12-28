package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// GetListIntent used to get a list of tasks from the repository.
type GetListIntent struct {
	ListRepository ListRepository
}

// Enact takes list id as URL parameter and serves a list of tasks.
func (i GetListIntent) Enact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mocklist, err := i.ListRepository.Get(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(mocklist)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}
