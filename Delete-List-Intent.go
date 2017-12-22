package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//DeleteListIntent ...
type DeleteListIntent struct {
	ListRepository ListRepository
	TaskRepository TaskRepository
}

//Enact ...
func (i DeleteListIntent) Enact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := i.ListRepository.Delete(vars["id"])
	if err != nil {
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}
	err = i.TaskRepository.Deletelist(vars["id"])
	if err == nil {
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusBadRequest)

	} else {
		fmt.Fprintf(w, "delete list succesfully")
		w.WriteHeader(http.StatusOK)

	}
}
