package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	ID        int
	Name      string
	CreatedAt int
	Status    bool
}

type List struct {
	Name      string
	Tasks     []Task
	CreatedAt int
}

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

	lane()
	fmt.Println("ads")
	log.Fatal(http.ListenAndServe(":8000", router))

}

func lane() {
	database, _ := sql.Open("sqlite3", "./test.db")
	//statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS lists (listname CHAR(50) PRIMARY KEY, createdat CHAR(50))")

	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS tasks (ID INTEGER PRIMARY KEY AUTOINCREMENT, createdat CHAR(50),name CHAR(50),status bool,listname CHAR(50),FOREIGN KEY(listname) REFERENCES lists(listname))")
	statement.Exec()

}
