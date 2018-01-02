package main

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

// ListModel ...
type ListModel struct {
	listName  string
	createdAt string
}

// TaskModel ...
type TaskModel struct {
	ID        int
	name      string
	status    bool
	listName  string
	createdAt string
}

// ListRepository interface...
type ListRepository interface {
	Create(list List) error
	Delete(listName string) error
	Check(listName string) bool
	Get(listName string) (List, error)
	AddTaskToList(list List) error
	DeleteTask(ID string, listName string) error
	GetTask(ID string, listName string) (Task, error)
	UpdateTask(listName string, taskID string, state bool) error
}

// SQLiteListRepository implementing ListRepository interface.
type SQLiteListRepository struct {
}

// Get method to retrieve a list of tasks from database.
func (r SQLiteListRepository) Get(listName string) (List, error) {
	var (
		name      string
		createdAt string
		tasklist  []Task
		ID        int
		status    int
		Listname  string
		list      List
	)
	database, err := sql.Open(DbType, DbName)
	if err != nil {
		return List{}, err
	}
	defer database.Close()
	query := "SELECT * FROM lists where listname='" + listName + "'; "
	rows, err := database.Query(query)
	defer rows.Close()
	if err != nil {
		return List{}, err
	}
	count := 0
	for rows.Next() {
		rows.Scan(&name, &createdAt)
		count++
	}
	fmt.Println(count)
	if count == 0 {
		return List{}, errors.New("List doesnt exist")
	}
	list.Name = name
	list.CreatedAt = createdAt
	query = "SELECT * FROM tasks where listname='" + listName + "'; "
	rows, err = database.Query(query)

	for rows.Next() {
		rows.Scan(&ID, &createdAt, &name, &status, &Listname)
		task := Task{ID: ID, Name: name, CreatedAt: createdAt}
		if status == 0 {
			task.Status = false
		} else {
			task.Status = true
		}
		tasklist = append(tasklist, task)
	}
	list.Tasks = tasklist
	return list, err
}

// Create method adds a list to the repository.
func (r SQLiteListRepository) Create(list List) error {
	if r.Check(list.Name) == true {
		return errors.New("List already exist")
	}
	listmodel := r.listModelFactory(list, time.Now().Local().Format(TimeFormat))
	database, err := sql.Open(DbType, DbName)
	if err != nil {
		return err
	}
	defer database.Close()
	query := "INSERT INTO lists (listname , createdat) VALUES ( '" +
		listmodel.listName + "' , '" + listmodel.createdAt + "')"
	statement, err := database.Prepare(query)
	statement.Exec()
	return err
}

// Delete method deletes list from the database.
func (r SQLiteListRepository) Delete(listName string) error {
	if r.Check(listName) == false {
		return errors.New("List not found")
	}
	database, err := sql.Open(DbType, DbName)
	if err != nil {
		return err
	}
	defer database.Close()
	query := "DELETE FROM tasks WHERE listname = '" + listName + "';"
	statement, err := database.Prepare(query)
	statement.Exec()
	if err != nil {
		return err
	}
	query = "DELETE FROM lists WHERE listname = '" + listName + "';"
	statement, err = database.Prepare(query)
	statement.Exec()
	return err
}

// Check method checks if a list exists in the database.
func (r SQLiteListRepository) Check(listName string) bool {
	database, err := sql.Open(DbType, DbName)
	if err != nil {
		return false
	}
	defer database.Close()
	query := "SELECT listname FROM lists where listname='" + listName + "'; "
	rows, err := database.Query(query)
	defer rows.Close()
	if err != nil {
		return false
	}
	return rows.Next()
}

// DeleteTask method deletes a task from the database.
func (r SQLiteListRepository) DeleteTask(ID string, listName string) error {
	if r.Check(listName) == false {
		return errors.New("List doesnt exist")
	}
	database, err := sql.Open(DbType, DbName)
	if err != nil {
		return err
	}
	defer database.Close()
	query := "DELETE FROM tasks WHERE ID = '" + ID + "' ;"
	statement, err := database.Prepare(query)
	res, err := statement.Exec()
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if affected == 0 {
		return errors.New("Entry Not found")
	}
	return err
}

// AddTaskToList method adds task to the database.
func (r SQLiteListRepository) AddTaskToList(list List) error {
	check := r.Check(list.Name)
	if check == false {
		return errors.New("List doesnt exists")
	}
	createdAt := time.Now().Local().Format(TimeFormat)
	task := taskModelFactory(list.Tasks[0], list.Name, createdAt)
	database, err := sql.Open(DbType, DbName) //enviroment variables
	if err != nil {
		return err
	}
	defer database.Close()
	query := "INSERT INTO tasks (createdat,name,status,listname) VALUES ( '" +
		task.createdAt + "','" + task.name + "','0','" + task.listName + "')"
	statement, err := database.Prepare(query)
	statement.Exec()
	return err
}

// GetTask method retrieves a task from the database.
func (r SQLiteListRepository) GetTask(ID string, list string) (Task, error) {
	var (
		task      Task
		name      string
		id        int
		createdAt string
		status    int
		listName  string
	)
	if r.Check(list) == false {
		return Task{}, errors.New("List doesnt exists")
	}
	database, err := sql.Open(DbType, DbName)
	if err != nil {
		return Task{}, err
	}
	defer database.Close()
	query := "SELECT * FROM tasks WHERE ID = '" + ID + "' ;"
	rows, err := database.Query(query)
	if err != nil {
		return Task{}, err
	}
	for rows.Next() {
		rows.Scan(&id, &createdAt, &name, &status, &listName)
		task.ID = id
		task.CreatedAt = createdAt
		task.Name = name
		if status == 0 {
			task.Status = false
		} else {
			task.Status = true
		}
	}
	return task, nil
}

// UpdateTask method updates task status.
func (r SQLiteListRepository) UpdateTask(listName string, taskID string, state bool) error {
	var status string
	if r.Check(listName) == false {
		return errors.New("List not found")
	}
	database, err := sql.Open(DbType, DbName)
	if err != nil {
		return err
	}
	defer database.Close()
	if state == false {
		status = "0"
	} else {
		status = "1"
	}
	query := "UPDATE tasks SET status= '" + status +
		"' WHERE ID= '" + taskID + "';"
	statement, err := database.Prepare(query)
	res, err := statement.Exec()
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if affected == 0 {
		return errors.New("Updation failed")
	}
	return err
}

func taskModelFactory(tasks Task, listName string, createdAt string) TaskModel {
	taskmodel := TaskModel{name: tasks.Name, status: false, listName: listName, createdAt: createdAt}
	return taskmodel
}

func (r SQLiteListRepository) listModelFactory(list List, time string) ListModel {
	listmodel := ListModel{listName: list.Name, createdAt: time}
	return listmodel
}
