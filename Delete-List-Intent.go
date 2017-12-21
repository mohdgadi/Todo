package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//DeleteListIntent ...
type DeleteListIntent struct {
	ListRepository ListRepository
}

//Enact ...
func (i DeleteListIntent) Enact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	val, err := i.ListRepository.Delete(vars["id"])
	if err == nil && val != 0 {
		fmt.Fprintf(w, "delete list succesfully")
	} else {
		fmt.Fprintf(w, "delete list unsuccesfully")
	}
}
