package main

import (
	"database/sql"
	"fmt"
	"time"
)

//ListModel ...
type ListModel struct {
	listName  string
	createdAt string
}

//ListRepository ...
type ListRepository interface {
	Create(list List) (int, error)
	Delete(listName string) (int, error)
	Check(listname string) bool
	Get(listname string) (List, error)
}

//SQLiteListRepository ...
type SQLiteListRepository struct {
}

//Get ...
func (r SQLiteListRepository) Get(listname string) (List, error) {
	database, _ := sql.Open("sqlite3", "./test.db")
	defer database.Close()
	query := "SELECT * FROM lists where listname='" + listname + "'; "
	rows, err := database.Query(query)
	var list = List{}
	if err != nil {
		return list, err
	}
	var name string
	var createdat string
	count := 0
	for rows.Next() {
		count++
		rows.Scan(&name, &createdat)

	}

	list.CreatedAt = createdat
	list.Name = name

	taskRepository := SQLiteTaskRepository{}
	list.Tasks, err = taskRepository.GetAll(listname)
	return list, nil

}

//Create ...
func (r SQLiteListRepository) Create(list List) (int, error) {
	if r.Check(list.Name) == true {
		fmt.Println("list already exists")
		return 0, nil
	} else {
		fmt.Println("list not exists")
	}
	listmodel := ListModel{listName: list.Name, createdAt: time.Now().Local().Format("2006-01-02")}
	database, _ := sql.Open("sqlite3", "./test.db")
	defer database.Close()
	var query = "INSERT INTO lists (listname , createdat) VALUES ( '" + listmodel.listName + "' , '" + listmodel.createdAt + "')"
	statement, err := database.Prepare(query)
	statement.Exec()
	if err != nil {
		return 0, err
	}
	return 1, nil
}

//Delete ..
func (r SQLiteListRepository) Delete(listname string) (int, error) {
	if r.Check(listname) == false {

		fmt.Println("List doesnt exist of the name")
		return 0, nil
	}
	database, _ := sql.Open("sqlite3", "./test.db")
	defer database.Close()
	var query = "DELETE FROM lists WHERE listname = '" + listname + "';"
	statement, err := database.Prepare(query)
	statement.Exec()
	if err != nil {
		return 0, err
	}
	taskRepository := SQLiteTaskRepository{}
	_, err = taskRepository.Deletelist(listname)
	if err != nil {
		return 0, err
	}
	return 1, nil

}

//Check ...
func (r SQLiteListRepository) Check(listname string) bool {
	database, _ := sql.Open("sqlite3", "./test.db")
	defer database.Close()
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
	}
	return true

}
