package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router := mux.NewRouter()

	getlistintent := GetListIntent{}
	createlistintent := CreateListIntent{}
	deletelistintent := DeleteListIntent{}
	deletetaskintent := DeleteTaskIntent{}
	updatetaskintent := UpdateTaskIntent{}
	gettaskintent := GetTaskIntent{}
	addtaskintent := AddTaskIntent{}

	router.HandleFunc("/lists", createlistintent.Enact).Methods("POST")
	router.HandleFunc("/lists/{id}", deletelistintent.Enact).Methods("DELETE")
	router.HandleFunc("/lists/{id}", getlistintent.Enact).Methods("GET")
	router.HandleFunc("/tasks/{id}", gettaskintent.Enact).Methods("GET")
	router.HandleFunc("/tasks", addtaskintent.Enact).Methods("POST")
	router.HandleFunc("/tasks/{id}", deletetaskintent.Enact).Methods("DELETE")
	router.HandleFunc("/tasks", updatetaskintent.Enact).Methods("PUT")
	//Taskrepo := TaskRepo{}

	//task := TaskModel{name: "task732327", listName: "list3", createdAt: t}
	//Taskrepo.addtask(task)
	log.Fatal(http.ListenAndServe(":8000", router))
	//listRepo = NewListRepo("myfile.db")
	//createlistintent := CreeateListIntent{listRepo}

}
