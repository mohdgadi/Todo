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

	mocklist := Task{
		Name:      "Task 1",
		CreatedAt: 2200,
		ID:        1,
		Status:    true,
	}
	data, err := json.Marshal(mocklist)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	fmt.Println("task", vars["id"], " Delivered successfuly")
}
