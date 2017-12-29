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
	getListIntent := GetListIntent{listRepository}
	createListIntent := CreateListIntent{listRepository}
	deleteListIntent := DeleteListIntent{listRepository}
	deleteTaskIntent := DeleteTaskIntent{listRepository}
	updateTaskIntent := UpdateTaskIntent{listRepository}
	getTaskIntent := GetTaskIntent{listRepository}
	addTaskIntent := AddTaskIntent{listRepository}

	router.HandleFunc("/lists", createListIntent.Enact).Methods("POST")
	router.HandleFunc("/lists/{id}", deleteListIntent.Enact).Methods("DELETE")
	router.HandleFunc("/lists/{id}", getListIntent.Enact).Methods("GET")
	router.HandleFunc("/lists/{listid}/tasks/{taskid}", getTaskIntent.Enact).Methods("GET")
	router.HandleFunc("/lists/{listid}/tasks", addTaskIntent.Enact).Methods("POST")
	router.HandleFunc("/lists/{listid}/tasks/{taskid}", deleteTaskIntent.Enact).Methods("DELETE")
	router.HandleFunc("/lists/{listid}/tasks/{taskid}", updateTaskIntent.Enact).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", router))
}
