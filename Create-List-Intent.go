package main

import (
	"encoding/json"
	"fmt"

	"net/http"
)

// CreateListIntent used to create list in the repository.
type CreateListIntent struct {
	ListRepository ListRepository
}

// Enact takes JSON request and creates a list.
func (c CreateListIntent) Enact(w http.ResponseWriter, r *http.Request) {
	var reqJSON List
	if err := json.NewDecoder(r.Body).Decode(&reqJSON); err == nil {
		if reqJSON.Name == "" {
			http.Error(w, "Request name cant be empty", http.StatusBadRequest)
		}
		list := List{Name: reqJSON.Name}
		err := c.ListRepository.Create(list)
		if err == nil {
			fmt.Fprintf(w, "created successfully")
			w.WriteHeader(http.StatusOK)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
