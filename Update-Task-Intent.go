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

	taskrepo := TaskRepo{}
	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusBadRequest)
	} else {

		if err = json.Unmarshal(data, &reqJSON); err == nil {
			val, err := taskrepo.update(reqJSON)
			if err == nil && val != 0 {
				fmt.Fprintf(w, "Task updated succesfully")
			} else {
				fmt.Fprintf(w, "TAsk addition unsuccesful")
				w.WriteHeader(http.StatusBadRequest)
			}
		} else {
			fmt.Fprintf(w, err.Error())
			w.WriteHeader(http.StatusBadRequest)
		}
	}
	fmt.Println("Update Task", reqJSON.ID, " to ", reqJSON.Status)
	fmt.Fprintf(w, "Updated task status")
}
