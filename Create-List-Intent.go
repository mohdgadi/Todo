package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"
)

//CreateListIntent ...
type CreateListIntent struct {
	ListRepository ListRepository
}

//Enact ...
func (c CreateListIntent) Enact(w http.ResponseWriter, r *http.Request) {
	var reqJSON List
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusBadRequest)
	} else {
		if err = json.Unmarshal(data, &reqJSON); err == nil {
			if reqJSON.Name == "" {
				fmt.Fprintf(w, "Request Name cant be empty")
				w.WriteHeader(http.StatusBadRequest)
			}
			list := List{Name: reqJSON.Name}
			err := c.ListRepository.Create(list)
			if err == nil {
				fmt.Fprintf(w, "created successfully")
				fmt.Println("created successfully")
			} else {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, err.Error())
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
		}
	}
}
