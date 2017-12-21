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
	} else {

		if err = json.Unmarshal(data, &reqJSON); err == nil {
			list := List{Name: reqJSON.Name}
			val, err := c.ListRepository.Create(list)
			if err == nil && val != 0 {
				fmt.Fprintf(w, "created successfully")
				fmt.Println("created successfully")
			} else {
				fmt.Fprintf(w, "Error occured")
				w.WriteHeader(http.StatusBadRequest)
			}

		} else {
			fmt.Fprintf(w, "Error occured")

			w.WriteHeader(http.StatusBadRequest)
		}
	}

}
