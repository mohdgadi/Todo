package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	ID        int
	Name      string
	CreatedAt string
	Status    bool
}

type List struct {
	Name      string
	Tasks     []Task
	CreatedAt string
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
	//t := time.Now().Local().Format("2006-01-02")

	//list := ListModel{listName: "list 2", createdAt: t}
	//task := TaskModel{name: "task7327", listName: "list3", createdAt: t}
	//addtask(task)
	//createlist(list)
	//deletelist("list 2")
	//deletetask(2)
	//gettask(1)
	getlist("list99")
	log.Fatal(http.ListenAndServe(":8000", router))
	//listRepo = NewListRepo("myfile.db")
	//createlistintent := CreeateListIntent{listRepo}

}

func createlist(list ListModel) {
	database, _ := sql.Open("sqlite3", "./test.db")
	var query = "INSERT INTO lists (listname , createdat) VALUES ( '  " + list.listName + " ' , ' " + list.createdAt + " ' )"
	fmt.Println(query)
	statement, _ := database.Prepare(query)
	statement.Exec()

}
func deletelist(listname string) {
	database, _ := sql.Open("sqlite3", "./test.db")
	var query = "DELETE FROM lists WHERE listname = '" + listname + "' ;"
	statement, err := database.Prepare(query)
	statement.Exec()
	if err != nil {
		panic(err)
	}

}
func addtask(task TaskModel) {
	database, _ := sql.Open("sqlite3", "./test.db")
	if checklist(task.listName) {
		var query = "INSERT INTO tasks (createdat,name,status,listname) VALUES ( '" + task.createdAt + "','" + task.name + "','0','" + task.listName + "')"
		statement, err := database.Prepare(query)
		if err != nil {
			panic(err)
		}
		statement.Exec()
	} else {
		fmt.Println("error occured")
	}

}
func checklist(listname string) bool {
	fmt.Println("ads")
	database, _ := sql.Open("sqlite3", "./test.db")
	query := "SELECT listname FROM lists where listname='" + listname + "'; "
	rows, err := database.Query(query)
	if err != nil {
		panic(err)
		return false
	}
	var name string
	count := 0
	for rows.Next() {
		count++
		rows.Scan(&name)

	}
	if count == 0 {
		return false
	} else {
		return true
	}
}
func deletetask(ID int) {
	database, _ := sql.Open("sqlite3", "./test.db")
	var query = "DELETE FROM tasks WHERE ID = '" + strconv.Itoa(ID) + "' ;"
	statement, err := database.Prepare(query)
	statement.Exec()
	if err != nil {
		panic(err)
	}
}
func gettask(ID int) {
	database, _ := sql.Open("sqlite3", "./test.db")
	var query = "SELECT * FROM tasks WHERE ID = '" + strconv.Itoa(ID) + "' ;"
	rows, err := database.Query(query)
	if err != nil {
		panic(err)
	}
	var task Task
	var name string
	var Id int
	var createdat string
	var status int
	var listname string
	for rows.Next() {
		rows.Scan(&Id, &createdat, &name, &status, &listname)
		task.ID = Id
		task.CreatedAt = createdat
		task.Name = name
		if status == 0 {
			task.Status = false
		} else {
			task.Status = true
		}
		fmt.Println(strconv.Itoa(Id) + ": " + createdat + " " + name + " " + listname + " " + strconv.Itoa(status))

	}
}

func getlist(listname string) {
	database, _ := sql.Open("sqlite3", "./test.db")

	query := "SELECT * FROM lists where listname='" + listname + "'; "
	rows, err := database.Query(query)
	if err != nil {
		panic(err)

	}
	var name string
	var createdat string
	var list = List{}
	count := 0
	for rows.Next() {
		count++
		rows.Scan(&name, &createdat)

	}

	list.CreatedAt = createdat
	list.Name = name

	query = "SELECT * FROM tasks where listname='" + listname + "'; "
	rows, err = database.Query(query)
	if err != nil {
		panic(err)

	}

	var tasklist []Task
	var Id int
	var status int
	var Listname string
	for rows.Next() {
		rows.Scan(&Id, &createdat, &name, &status, &Listname)
		task := Task{ID: Id, Name: name, CreatedAt: createdat}
		if status == 0 {
			task.Status = false
		} else {
			task.Status = true
		}
		tasklist = append(tasklist, task)

	}
	list.Tasks = tasklist
	fmt.Println(list)

}

/**/
