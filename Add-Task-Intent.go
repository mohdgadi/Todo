package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"
)

type AddTaskIntent struct {
}

func (AddTaskIntent) Enact(w http.ResponseWriter, r *http.Request) {

	var reqJSON List
	var task Task
	taskrepo := TaskRepo{}
	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusBadRequest)
	} else {

		if err = json.Unmarshal(data, &reqJSON); err == nil {
			task = reqJSON.Tasks[0]
			val, err := taskrepo.add(task, reqJSON.Name)
			if err == nil && val != 0 {
				fmt.Fprintf(w, "Task added succesfully")
			} else {
				fmt.Fprintf(w, "TAsk addition unsuccesful")
			}

		} else {
			fmt.Fprintf(w, err.Error())
			w.WriteHeader(http.StatusBadRequest)
		}
	}

}
