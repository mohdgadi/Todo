package main

import (
	"database/sql"
	"errors"
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
	Create(list List) error
	Delete(listName string) error
	Check(listname string) bool
	Get(listname string) (List, error)
	ListModelFactory(list List, time string) ListModel
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
	defer rows.Close() //error after this

	if err != nil {
		return List{}, err
	}
	list := List{}
	var name string
	var createdat string
	count := 0 // Code cleanup
	for rows.Next() {
		count++
		rows.Scan(&name, &createdat)

	}
	list.CreatedAt = createdat
	list.Name = name
	return list, nil

}

//Create ...
func (r SQLiteListRepository) Create(list List) error {
	if r.Check(list.Name) == true {
		fmt.Println("list already exists")
		return errors.New("List already exist")
	}
	//factory
	listmodel := r.ListModelFactory(list, time.Now().Local().Format("2006-01-02"))
	database, _ := sql.Open("sqlite3", "./test.db")
	defer database.Close()
	var query = "INSERT INTO lists (listname , createdat) VALUES ( '" + listmodel.listName + "' , '" + listmodel.createdAt + "')"
	statement, err := database.Prepare(query)
	statement.Exec()
	if err != nil {
		return err // use custom error
	}
	return nil
}

//Delete ..
func (r SQLiteListRepository) Delete(listname string) error {
	if r.Check(listname) == false {

		fmt.Println("List doesnt exist of the name")
		return errors.New("List not found")
	}
	database, _ := sql.Open("sqlite3", "./test.db")
	defer database.Close()
	var query = "DELETE FROM lists WHERE listname = '" + listname + "';"
	statement, err := database.Prepare(query)
	statement.Exec()
	if err != nil {
		return err
	}
	return nil
}

//Check ...
func (r SQLiteListRepository) Check(listname string) bool {
	database, _ := sql.Open("sqlite3", "./test.db")
	defer database.Close()
	query := "SELECT listname FROM lists where listname='" + listname + "'; "
	rows, err := database.Query(query)
	defer rows.Close()
	if err != nil {
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

//ListModelFactory ...
func (r SQLiteListRepository) ListModelFactory(list List, time string) ListModel {
	listmodel := ListModel{listName: list.Name, createdAt: time}
	return listmodel
}
