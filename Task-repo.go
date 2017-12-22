package main

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"
)

//TaskModel ...
type TaskModel struct {
	ID        int
	name      string
	status    bool
	listName  string
	createdAt string
}

// TaskRepository ...
type TaskRepository interface {
	Add(task Task, listname string) error
	Delete(ID string) error
	Deletelist(listname string) error
	Get(ID string) (Task, error)
	GetAll(listname string) ([]Task, error)
	Update(task Task) error
}

//SQLiteTaskRepository ...
type SQLiteTaskRepository struct {
}

//Add ...
func (t SQLiteTaskRepository) Add(tasks Task, listname string) error {

	task := TaskModel{name: tasks.Name, status: false, listName: listname} //use factory
	task.createdAt = time.Now().Local().Format("2006-01-02")

	database, _ := sql.Open("sqlite3", "./test.db") //enviroment variables
	defer database.Close()

	var query = "INSERT INTO tasks (createdat,name,status,listname) VALUES ( '" + task.createdAt + "','" + task.name + "','0','" + task.listName + "')"
	statement, err := database.Prepare(query)
	statement.Exec()
	if err != nil {
		return err
	}
	return nil
}

//Delete ...
func (t SQLiteTaskRepository) Delete(ID string) error {
	database, _ := sql.Open("sqlite3", "./test.db")
	defer database.Close()
	var query = "DELETE FROM tasks WHERE ID = '" + ID + "' ;"
	statement, err := database.Prepare(query)
	res, err := statement.Exec()
	affected, err := res.RowsAffected()
	if affected == 0 {
		return errors.New("Entry Not found")
	}
	if err != nil {
		return err
	}
	return nil
}

//Get ...
func (t SQLiteTaskRepository) Get(ID string) (Task, error) {
	database, _ := sql.Open("sqlite3", "./test.db")
	defer database.Close()
	var query = "SELECT * FROM tasks WHERE ID = '" + ID + "' ;"
	rows, err := database.Query(query)
	var task Task
	if err != nil {
		return task, err
	}
	var name string
	var id int
	var createdat string
	var status int
	var listname string
	for rows.Next() {
		rows.Scan(&id, &createdat, &name, &status, &listname)
		task.ID = id
		task.CreatedAt = createdat
		task.Name = name
		if status == 0 {
			task.Status = false
		} else {
			task.Status = true
		}
	}
	return task, nil
}

//Deletelist ...
func (t SQLiteTaskRepository) Deletelist(listname string) error {
	database, _ := sql.Open("sqlite3", "./test.db")
	defer database.Close()
	var query = "DELETE FROM tasks WHERE listname = '" + listname + "';"
	statement, err := database.Prepare(query)
	statement.Exec()
	if err != nil {
		return err
	}
	fmt.Println("deleted task list")
	return nil
}

//Update ...
func (t SQLiteTaskRepository) Update(task Task) error {
	database, _ := sql.Open("sqlite3", "./test.db")
	defer database.Close()
	var status string
	if task.Status == false {
		status = "0"
	} else {
		status = "1"
	}
	var query = "UPDATE tasks SET status= '" + status + "' WHERE ID= '" + strconv.Itoa(task.ID) + "';"
	statement, err := database.Prepare(query)
	_, err = statement.Exec()

	if err != nil {
		return err
	}
	return nil
}

//GetAll ...
func (t SQLiteTaskRepository) GetAll(listname string) ([]Task, error) {
	database, _ := sql.Open("sqlite3", "./test.db")
	defer database.Close()
	query := "SELECT * FROM tasks where listname='" + listname + "'; "
	rows, err := database.Query(query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var createdat string
	var tasklist []Task
	var id int
	var status int
	var Listname string
	var name string
	for rows.Next() {
		rows.Scan(&id, &createdat, &name, &status, &Listname)
		task := Task{ID: id, Name: name, CreatedAt: createdat}
		if status == 0 {
			task.Status = false
		} else {
			task.Status = true
		}
		tasklist = append(tasklist, task)

	}
	return tasklist, nil
}
