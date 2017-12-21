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
}

//Enact ...
func (i GetListIntent) Enact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	mocklist, err := i.ListRepository.Get(vars["id"])
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
