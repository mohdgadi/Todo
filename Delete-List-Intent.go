package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type DeleteListIntent struct {
}

func (d DeleteListIntent) Enact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	listrepo := ListRepo{}

	val, err := listrepo.delete(vars["id"])
	if err == nil && val != 0 {
		fmt.Fprintf(w, "delete list succesfully")
	} else {
		fmt.Fprintf(w, "delete list unsuccesfully")
	}
}
