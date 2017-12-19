package main

import (
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
)

type DeleteTaskIntent struct {
}

func (DeleteTaskIntent) Enact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["id"], " task deleted succesfully")
	fmt.Fprintf(w, "delete task succesfully")
}
