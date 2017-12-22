package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//GetListIntent ...
type GetListIntent struct {
	ListRepository ListRepository
	TaskRepository TaskRepository
}

//Enact ...
func (i GetListIntent) Enact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mocklist, err := i.ListRepository.Get(vars["id"])
	if err != nil {
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}
	mocklist.Tasks, err = i.TaskRepository.GetAll(vars["id"])
	if err != nil {
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}
	// Error handling shouldn't be neglected. Can use OKGO.
	data, err := json.Marshal(mocklist)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	fmt.Println("list ", vars["id"], " Delivered successfuly")
}
