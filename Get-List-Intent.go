package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type GetListIntent struct {
}

func (GetListIntent) Enact(w http.ResponseWriter, r *http.Request) {
	listrepo := ListRepo{}
	vars := mux.Vars(r)

	mocklist, err := listrepo.get(vars["id"])
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
