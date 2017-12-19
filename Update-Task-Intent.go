package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"
)

type UpdateTaskIntent struct {
}

func (UpdateTaskIntent) Enact(w http.ResponseWriter, r *http.Request) {
	var reqJSON Task

	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {

		if err = json.Unmarshal(data, &reqJSON); err == nil {

			fmt.Fprintf(w, reqJSON.Name)
		} else {
			fmt.Fprintf(w, err.Error())
		}
	}
	fmt.Println("Update Task", reqJSON.ID, " to ", reqJSON.Status)
	fmt.Fprintf(w, "Updated task status")
}
