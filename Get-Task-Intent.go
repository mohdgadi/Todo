package main

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
)

type GetTaskIntent struct {
}

func (GetTaskIntent) Enact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskrepo := TaskRepo{}
	mocklist, err := taskrepo.get(vars["id"])
	if mocklist.ID == 0 || err != nil {
		fmt.Println("Doestn exist")
		w.WriteHeader(http.StatusBadRequest)
	} else {
		data, err := json.Marshal(mocklist)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Bad request")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
		fmt.Println("task", vars["id"], " Delivered successfuly")
	}

}
