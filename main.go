package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router := mux.NewRouter()
	listRepository := SQLiteListRepository{}
	taskRepository := SQLiteTaskRepository{}
	getlistintent := GetListIntent{listRepository}
	createlistintent := CreateListIntent{listRepository}
	deletelistintent := DeleteListIntent{listRepository}
	deletetaskintent := DeleteTaskIntent{taskRepository}
	updatetaskintent := UpdateTaskIntent{taskRepository}
	gettaskintent := GetTaskIntent{taskRepository}
	addtaskintent := AddTaskIntent{taskRepository}

	router.HandleFunc("/lists", createlistintent.Enact).Methods("POST")
	router.HandleFunc("/lists/{id}", deletelistintent.Enact).Methods("DELETE")
	router.HandleFunc("/lists/{id}", getlistintent.Enact).Methods("GET")
	router.HandleFunc("/tasks/{id}", gettaskintent.Enact).Methods("GET")
	router.HandleFunc("/tasks", addtaskintent.Enact).Methods("POST")
	router.HandleFunc("/tasks/{id}", deletetaskintent.Enact).Methods("DELETE")
	router.HandleFunc("/tasks", updatetaskintent.Enact).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", router))

}
