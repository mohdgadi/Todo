package main

import (
	"database/sql"
	"fmt"
	"time"
)

type ListModel struct {
	listName  string
	createdAt string
}

type ListRep interface {
	create(list List) (int, error)
	delete(listName string) (int, error)
	check(listname string) bool
	get(listname string) (List, error)
}

type ListRepo struct {
}

func (r ListRepo) get(listname string) (List, error) {
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

	query = "SELECT * FROM tasks where listname='" + listname + "'; "
	rows, err = database.Query(query)
	if err != nil {
		return list, err

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
	return list, nil

}

func (r ListRepo) create(list List) (int, error) {
	if r.check(list.Name) == true {
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
func (r ListRepo) delete(listname string) (int, error) {
	if r.check(listname) == false {

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
	taskrepo := TaskRepo{}
	_, err = taskrepo.deletelist(listname)
	if err != nil {
		return 0, err
	}
	return 1, nil

}
func (r ListRepo) check(listname string) bool {
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
