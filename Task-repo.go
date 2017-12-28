package main

import (
	"database/sql"
	"errors"
	"strconv"
	"time"
)

// TaskModel ...
type TaskModel struct {
	ID        int
	name      string
	status    bool
	listName  string
	createdAt string
}

// TaskRepository interface.
type TaskRepository interface {
	Add(task Task, listName string) error
	Delete(ID string) error
	DeleteTaskList(listName string) error
	Get(ID string) (Task, error)
	GetTaskList(listName string) ([]Task, error)
	Update(task Task) error
}

// SQLiteTaskRepository implementing TaskRepository interface..
type SQLiteTaskRepository struct {
}

// Add method adds task to the database.
func (t SQLiteTaskRepository) Add(tasks Task, listName string) error {
	createdAt := time.Now().Local().Format(TimeFormat)
	task := TaskModelFactory(tasks, listName, createdAt)
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

// Delete method deletes a task from the database.
func (t SQLiteTaskRepository) Delete(ID string) error {
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

// Get method retrieves a task from the database.
func (t SQLiteTaskRepository) Get(ID string) (Task, error) {
	var (
		task      Task
		name      string
		id        int
		createdAt string
		status    int
		listName  string
	)
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

// DeleteTaskList method deletes a list of task having same listname.
func (t SQLiteTaskRepository) DeleteTaskList(listName string) error {
	database, err := sql.Open(DbType, DbName)
	if err != nil {
		return err
	}
	defer database.Close()
	query := "DELETE FROM tasks WHERE listname = '" + listName + "';"
	statement, err := database.Prepare(query)
	statement.Exec()
	return err
}

// Update method updates task status.
func (t SQLiteTaskRepository) Update(task Task) error {
	var status string
	database, err := sql.Open(DbType, DbName)
	if err != nil {
		return err
	}
	defer database.Close()
	if task.Status == false {
		status = "0"
	} else {
		status = "1"
	}
	query := "UPDATE tasks SET status= '" + status +
		"' WHERE ID= '" + strconv.Itoa(task.ID) + "';"
	statement, err := database.Prepare(query)
	_, err = statement.Exec()
	return err
}

// GetTaskList retrieves a list of task having list name.
func (t SQLiteTaskRepository) GetTaskList(listName string) ([]Task, error) {
	var (
		createdAt string
		taskList  []Task
		id        int
		status    int
		ListName  string
		name      string
	)
	database, err := sql.Open(DbType, DbName)
	if err != nil {
		return nil, err
	}
	defer database.Close()
	query := "SELECT * FROM tasks where listname='" + listName + "'; "
	rows, err := database.Query(query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		rows.Scan(&id, &createdAt, &name, &status, &ListName)
		task := Task{ID: id, Name: name, CreatedAt: createdAt}
		if status == 0 {
			task.Status = false
		} else {
			task.Status = true
		}
		taskList = append(taskList, task)
	}
	return taskList, nil
}

// TaskModelFactory takes input as Task and returns a TaskModel .
func TaskModelFactory(tasks Task, listName string, createdAt string) TaskModel {
	taskmodel := TaskModel{name: tasks.Name, status: false, listName: listName, createdAt: createdAt}
	return taskmodel
}
