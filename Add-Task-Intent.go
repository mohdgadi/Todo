package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"
)

//AddTaskIntent ...
type AddTaskIntent struct {
	TaskRepository TaskRepository
	ListRepository ListRepository
}

//Enact ...
func (i AddTaskIntent) Enact(w http.ResponseWriter, r *http.Request) {
	var reqJSON List
	var task Task
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusBadRequest)
	} else {
		if err = json.Unmarshal(data, &reqJSON); err == nil {
			task = reqJSON.Tasks[0]
			check := i.ListRepository.Check(reqJSON.Name)
			if check == false {
				fmt.Fprintf(w, "List doesnt exist")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			err = i.TaskRepository.Add(task, reqJSON.Name)
			if err == nil {
				fmt.Fprintf(w, "Task added succesfully")
				w.WriteHeader(http.StatusOK)

			} else {
				fmt.Fprintf(w, "TAsk addition unsuccesful")
				fmt.Fprintf(w, err.Error())
				w.WriteHeader(http.StatusBadRequest)

			}
		} else {
			fmt.Fprintf(w, err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
