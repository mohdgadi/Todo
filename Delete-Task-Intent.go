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
	taskrepo := TaskRepo{}
	val, err := taskrepo.delete(vars["id"])
	if err == nil && val != 0 {
		fmt.Fprintf(w, "delete task succesfully")
		fmt.Println(err, val)
	} else {
		fmt.Fprintf(w, "delete task unsuccesfully")
		fmt.Println(err, val)
	}
}
