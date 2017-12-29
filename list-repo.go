package main

import (
	"database/sql"
	"errors"
	"time"
)

// ListModel ...
type ListModel struct {
	listName  string
	createdAt string
}

// ListModelFactory takes input as List and returns a ListModel.
func (r SQLiteListRepository) ListModelFactory(list List, time string) ListModel {
	listmodel := ListModel{listName: list.Name, createdAt: time}
	return listmodel
}

// ListRepository interface...
type ListRepository interface {
	Create(list List) error
	Delete(listName string) error
	Check(listName string) bool
	Get(listName string) (List, error)
}

// SQLiteListRepository implementing ListRepository interface.
type SQLiteListRepository struct {
	TaskRepository TaskRepository
}

// Get method to retrieve a list of tasks from database.
func (r SQLiteListRepository) Get(listName string) (List, error) {
	var (
		name      string
		createdAt string
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
	if count == 0 {
		return List{}, errors.New("List doesnt exist")
	}
	list := List{CreatedAt: createdAt, Name: name}
	list.Tasks, err = r.TaskRepository.GetTaskList(listName)
	return list, err
}

// Create method adds a list to the repository.
func (r SQLiteListRepository) Create(list List) error {
	if r.Check(list.Name) == true {
		return errors.New("List already exist")
	}
	listmodel := r.ListModelFactory(list, time.Now().Local().Format(TimeFormat))
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
	err = r.TaskRepository.DeleteTaskList(listName)
	if err != nil {
		return err
	}
	query := "DELETE FROM lists WHERE listname = '" + listName + "';"
	statement, err := database.Prepare(query)
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
