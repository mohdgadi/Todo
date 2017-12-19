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
	fmt.Println("Tasks added ", " Added succesfuly")
}
