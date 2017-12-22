package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"
)

//UpdateTaskIntent ...
type UpdateTaskIntent struct {
	TaskRepository TaskRepository
}

//Enact ...
func (i UpdateTaskIntent) Enact(w http.ResponseWriter, r *http.Request) {
	var reqJSON Task
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusBadRequest)
	} else {
		if err = json.Unmarshal(data, &reqJSON); err == nil {
			err := i.TaskRepository.Update(reqJSON)
			if err == nil {
				fmt.Fprintf(w, "Task updated succesfully")
			} else {
				fmt.Fprintf(w, "TAsk updation unsuccesful")
				w.WriteHeader(http.StatusBadRequest)
			}
		} else {
			fmt.Fprintf(w, err.Error())
			w.WriteHeader(http.StatusBadRequest)
		}
	}
	fmt.Println("Update Task", reqJSON.ID, " to ", reqJSON.Status)
}
