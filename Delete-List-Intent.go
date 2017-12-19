package main

import (
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
)

type DeleteListIntent struct {
}

func (DeleteListIntent) Enact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["id"], " list deleted succesfully")
	fmt.Fprintf(w, "delete list succesfully")
}
