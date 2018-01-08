package main

import (
	"database/sql"
	"errors"
	"strconv"
	"time"
)

// ListModel ...
type ListModel struct {
	listName  string
	createdAt string
}

// TaskModel ...
type TaskModel struct {
	id        int
	name      string
	status    bool
	listName  string
	createdAt string
}

// ListRepository interface...
type ListRepository interface {
	Create(list List) error
	Delete(listName string) error
	Get(listName string) (List, error)
	AddTaskToList(list List) error
	DeleteTaskFromList(taskID string) error
	GetTaskFromList(taskID string) (Task, error)
	UpdateTaskStatus(List) error
	CheckIfExists(listName string) bool
	GetAllTasksFromList(listName string) ([]Task, error)
}

// SQLiteListRepository implementing ListRepository interface.
type SQLiteListRepository struct {
}

// Get method to retrieve a list of tasks from database.
func (r SQLiteListRepository) Get(listName string) (List, error) {
	var (
		name      string
		createdAt string
		list      List
	)
	database, err := sql.Open(DbType, DbName)
	if err != nil {
		return List{}, err
	}
	defer database.Close()
	query := "SELECT * FROM " + ListTableName + " where listname='" + listName + "'; "
	rows, err := database.Query(query)
	if err != nil {
		return List{}, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&name, &createdAt)
		if err != nil {
			return List{}, err
		}
	}
	if name == "" {
		return List{}, errors.New("List doesnt exist")
	}
	list.Name = name
	time, err := time.Parse(TimeFormat, createdAt)
	if err != nil {
		return List{}, err
	}
	list.CreatedAt = time
	list.Tasks, err = r.GetAllTasksFromList(listName)
	return list, err
}

// Create method adds a list to the repository.
func (r SQLiteListRepository) Create(list List) error {
	err := r.createListTableIfNotExists()
	if err != nil {
		return err
	}
	database, err := sql.Open(DbType, DbName)
	if err != nil {
		return err
	}
	defer database.Close()
	listmodel := r.listModelFactory(list, time.Now().Local().Format(TimeFormat))
	query := "INSERT INTO " + ListTableName + " (listname , createdat) VALUES ( '" +
		listmodel.listName + "' , '" + listmodel.createdAt + "')"
	statement, err := database.Prepare(query)
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	return err
}

// Delete method deletes list from the database.
func (r SQLiteListRepository) Delete(listName string) error {
	database, err := sql.Open(DbType, DbName)
	if err != nil {
		return err
	}
	defer database.Close()
	err = r.deleteAllTasksFromList(listName)
	if err != nil {
		return err
	}
	query := "DELETE FROM " + ListTableName + " WHERE listname = '" + listName + "';"
	statement, err := database.Prepare(query)
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	return err
}

// CheckIfExists method checks if a list exists in the database.
func (r SQLiteListRepository) CheckIfExists(listName string) bool {
	database, err := sql.Open(DbType, DbName)
	if err != nil {
		return false
	}
	defer database.Close()
	query := "SELECT listname FROM " + ListTableName + " where listname='" + listName + "'; "
	rows, err := database.Query(query)
	if err != nil {
		return false
	}
	defer rows.Close()
	return rows.Next()
}

// DeleteTaskFromList method deletes a task from the database.
func (r SQLiteListRepository) DeleteTaskFromList(taskID string) error {
	database, err := sql.Open(DbType, DbName)
	if err != nil {
		return err
	}
	defer database.Close()
	query := "DELETE FROM " + TaskTableName + " WHERE ID = '" + taskID + "' ;"
	statement, err := database.Prepare(query)
	if err != nil {
		return err
	}
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
	err := r.createTaskTableIfNotExists()
	if err != nil {
		return err
	}
	database, err := sql.Open(DbType, DbName)
	if err != nil {
		return err
	}
	defer database.Close()
	list.Tasks[0].CreatedAt = time.Now()
	list.Tasks[0].Status = false
	task := taskModelFactory(list.Tasks[0], list.Name)
	query := "INSERT INTO " + TaskTableName + " (createdat,name,status,listname) VALUES ( '" +
		task.createdAt + "','" + task.name + "','0','" + task.listName + "')"
	statement, err := database.Prepare(query)
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	return err
}

// GetTaskFromList method retrieves a task from the database.
func (r SQLiteListRepository) GetTaskFromList(taskID string) (Task, error) {
	var (
		taskList []Task
	)
	database, err := sql.Open(DbType, DbName)
	if err != nil {
		return Task{}, err
	}
	defer database.Close()
	query := "SELECT * FROM " + TaskTableName + " WHERE ID = '" + taskID + "' ;"
	rows, err := database.Query(query)
	if err != nil {
		return Task{}, err
	}
	defer rows.Close()
	taskList, err = r.rowsToTaskFactory(rows)
	if err != nil {
		return Task{}, err
	}
	return taskList[0], nil
}

// UpdateTaskStatus method updates task status.
func (r SQLiteListRepository) UpdateTaskStatus(list List) error {
	var status string
	task := list.Tasks[0]
	database, err := sql.Open(DbType, DbName)
	if err != nil {
		return err
	}
	defer database.Close()
	switch task.Status {
	case true:
		status = strconv.Itoa(Completed)
	default:
		status = strconv.Itoa(Pending)
	}
	taskID := strconv.Itoa(task.ID)
	query := "UPDATE " + TaskTableName + " SET status= '" + status +
		"' WHERE ID= '" + taskID + "';"
	statement, err := database.Prepare(query)
	if err != nil {
		return err
	}
	res, err := statement.Exec()
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("Updation failed")
	}
	return err
}

// deleteAllTasksFromList deletes all those task having the listname.
func (r SQLiteListRepository) deleteAllTasksFromList(listName string) error {
	database, err := sql.Open(DbType, DbName)
	if err != nil {
		return err
	}
	defer database.Close()
	query := "DELETE FROM " + TaskTableName + " WHERE listname = '" + listName + "';"
	statement, err := database.Prepare(query)
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	return err
}

// GetAllTasksFromList fetches an array of tasks from a given list and returns it.
func (r SQLiteListRepository) GetAllTasksFromList(listName string) ([]Task, error) {
	database, err := sql.Open(DbType, DbName)
	if err != nil {
		return nil, err
	}
	defer database.Close()
	query := "SELECT * FROM " + TaskTableName + " where listname='" + listName + "'; "
	rows, err := database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return r.rowsToTaskFactory(rows)
}

func taskModelFactory(tasks Task, listName string) TaskModel {
	taskmodel := TaskModel{name: tasks.Name, status: tasks.Status, listName: listName,
		createdAt: tasks.CreatedAt.Format(TimeFormat)}
	return taskmodel
}

func (r SQLiteListRepository) listModelFactory(list List, time string) ListModel {
	listmodel := ListModel{listName: list.Name, createdAt: time}
	return listmodel
}

func (r SQLiteListRepository) createListTableIfNotExists() error {
	database, err := sql.Open(DbType, DbName)
	if err != nil {
		return err
	}
	defer database.Close()
	query := "CREATE TABLE IF NOT EXISTS " + ListTableName +
		" (listname CHAR(50) PRIMARY KEY, createdat CHAR(50));"
	statement, err := database.Prepare(query)
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	return err
}

func (r SQLiteListRepository) createTaskTableIfNotExists() error {
	database, err := sql.Open(DbType, DbName)
	if err != nil {
		return err
	}
	defer database.Close()
	query := "CREATE TABLE IF NOT EXISTS " + TaskTableName + " (ID INTEGER PRIMARY KEY AUTOINCREMENT," +
		"createdat CHAR(50),name CHAR(50),status int," +
		"listname CHAR(50),FOREIGN KEY(listname) REFERENCES lists(listname));"
	statement, err := database.Prepare(query)
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	return err
}

func (r SQLiteListRepository) rowsToTaskFactory(rows *sql.Rows) ([]Task, error) {
	var (
		task      Task
		taskList  []Task
		id        int
		createdAt string
		name      string
		status    int
		lstName   string
	)
	for rows.Next() {
		err := rows.Scan(&id, &createdAt, &name, &status, &lstName)
		if err != nil {
			return nil, err
		}
		time, err := time.Parse(TimeFormat, createdAt)
		if err != nil {
			return nil, err
		}
		task = Task{ID: id, Name: name, CreatedAt: time}
		if status == 0 {
			task.Status = false
		} else {
			task.Status = true
		}
		taskList = append(taskList, task)
	}
	return taskList, nil
}
