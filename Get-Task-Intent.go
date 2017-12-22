package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//GetTaskIntent ...
type GetTaskIntent struct {
	TaskRepository TaskRepository
}

//Enact ...
func (i GetTaskIntent) Enact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mocklist, err := i.TaskRepository.Get(vars["id"])
	if mocklist.ID == 0 || err != nil {
		fmt.Println("Doestn exist")
		w.WriteHeader(http.StatusBadRequest)
	} else {
		data, err := json.Marshal(mocklist)
		if err != nil {
			fmt.Fprint(w, "Bad request")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
		fmt.Println("task", vars["id"], " Delivered successfuly")
	}

}
