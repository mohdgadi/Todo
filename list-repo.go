package main

import (
	"database/sql"
	"errors"
	"strconv"
	"time"

	"github.com/angadn/okgo"
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
		database  *sql.DB
		rows      *sql.Rows
		times     time.Time
	)
	var err error
	okgo.NewOKGO(&err).On(func() error {
		database, err = sql.Open(DbType, DbName)
		return err
	}).On(func() error {
		query := "SELECT * FROM " + ListTableName + " where listname='" + listName + "'; "
		rows, err = database.Query(query)
		return err
	}).On(func() error {
		for rows.Next() {
			err = rows.Scan(&name, &createdAt)
			if err != nil {
				break
			}
		}
		return err
	}).On(func() error {
		if name == "" {
			err = errors.New("List doesnt exist")
		}
		return err
	}).On(func() error {
		list.Name = name
		times, err = time.Parse(TimeFormat, createdAt)
		return err
	}).On(func() error {
		list.CreatedAt = times
		list.Tasks, err = r.GetAllTasksFromList(listName)
		return err
	}).Run()
	if err != nil {
		return List{}, err
	}
	defer rows.Close()
	defer database.Close()
	return list, err
}

// Create method adds a list to the repository.
func (r SQLiteListRepository) Create(list List) error {
	var (
		err       error
		listmodel ListModel
		query     string
		statement *sql.Stmt
		database  *sql.DB
	)
	okgo.NewOKGO(&err).On(func() error {
		err = r.createListTableIfNotExists()
		return err
	}).On(func() error {
		database, err = sql.Open(DbType, DbName)
		return err
	}).On(func() error {
		listmodel = r.listModelFactory(list, time.Now().Local().Format(TimeFormat))
		query = "INSERT INTO " + ListTableName + " (listname , createdat) VALUES ( '" +
			listmodel.listName + "' , '" + listmodel.createdAt + "')"
		statement, err = database.Prepare(query)
		return err
	}).On(func() error {
		_, err = statement.Exec()
		return err
	}).Run()
	defer database.Close()
	return err
}

// Delete method deletes list from the database.
func (r SQLiteListRepository) Delete(listName string) error {
	var (
		err       error
		database  *sql.DB
		statement *sql.Stmt
	)
	okgo.NewOKGO(&err).On(func() error {
		database, err = sql.Open(DbType, DbName)
		return err
	}).On(func() error {
		err = r.deleteAllTasksFromList(listName)
		return err
	}).On(func() error {
		query := "DELETE FROM " + ListTableName + " WHERE listname = '" + listName + "';"
		statement, err = database.Prepare(query)
		return err
	}).On(func() error {
		_, err = statement.Exec()
		return err
	}).Run()
	defer database.Close()
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
	var (
		err       error
		database  *sql.DB
		statement *sql.Stmt
		affected  int64
		res       sql.Result
	)
	okgo.NewOKGO(&err).On(func() error {
		database, err = sql.Open(DbType, DbName)
		return err
	}).On(func() error {
		query := "DELETE FROM " + TaskTableName + " WHERE ID = '" + taskID + "' ;"
		statement, err = database.Prepare(query)
		return err
	}).On(func() error {
		res, err = statement.Exec()
		return err
	}).On(func() error {
		affected, err = res.RowsAffected()
		return err
	}).On(func() error {
		if affected == 0 {
			return errors.New("Entry Not found")
		}
		return nil
	}).Run()
	defer database.Close()
	return err
}

// AddTaskToList method adds task to the database.
func (r SQLiteListRepository) AddTaskToList(list List) error {
	var (
		err       error
		database  *sql.DB
		statement *sql.Stmt
	)
	okgo.NewOKGO(&err).On(func() error {
		err = r.createTaskTableIfNotExists()
		return err
	}).On(func() error {
		database, err = sql.Open(DbType, DbName)
		return err
	}).On(func() error {
		list.Tasks[0].CreatedAt = time.Now()
		list.Tasks[0].Status = false
		task := taskModelFactory(list.Tasks[0], list.Name)
		query := "INSERT INTO " + TaskTableName + " (createdat,name,status,listname) VALUES ( '" +
			task.createdAt + "','" + task.name + "','0','" + list.Name + "')"
		statement, err = database.Prepare(query)
		return err
	}).On(func() error {
		_, err = statement.Exec()
		return err
	}).Run()
	defer database.Close()
	return err
}

// GetTaskFromList method retrieves a task from the database.
func (r SQLiteListRepository) GetTaskFromList(taskID string) (Task, error) {
	var (
		err      error
		database *sql.DB
		taskList []Task
		rows     *sql.Rows
	)
	okgo.NewOKGO(&err).On(func() error {
		database, err = sql.Open(DbType, DbName)
		return err
	}).On(func() error {
		query := "SELECT * FROM " + TaskTableName + " WHERE ID = '" + taskID + "' ;"
		rows, err = database.Query(query)
		return err
	}).On(func() error {
		taskList, err = r.rowsToTaskFactory(rows)
		return err
	}).Run()
	if err != nil {
		return Task{}, err
	}
	defer rows.Close()
	defer database.Close()
	return taskList[0], nil
}

// UpdateTaskStatus method updates task status.
func (r SQLiteListRepository) UpdateTaskStatus(list List) error {
	var (
		err       error
		database  *sql.DB
		status    string
		task      Task
		res       sql.Result
		statement *sql.Stmt
		affected  int64
	)

	okgo.NewOKGO(&err).On(func() error {
		task = list.Tasks[0]
		database, err = sql.Open(DbType, DbName)
		return err
	}).On(func() error {
		switch task.Status {
		case true:
			status = strconv.Itoa(Completed)
		default:
			status = strconv.Itoa(Pending)
		}
		taskID := strconv.Itoa(task.ID)
		query := "UPDATE " + TaskTableName + " SET status= '" + status +
			"' WHERE ID= '" + taskID + "';"
		statement, err = database.Prepare(query)
		return err
	}).On(func() error {
		res, err = statement.Exec()
		return err
	}).On(func() error {
		affected, err = res.RowsAffected()
		return err
	}).On(func() error {
		if affected == 0 {
			return errors.New("Updation failed")
		}
		return nil
	}).Run()
	defer database.Close()
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
		err       error
		times     time.Time
	)
	for rows.Next() {
		okgo.NewOKGO(&err).On(func() error {
			err = rows.Scan(&id, &createdAt, &name, &status, &lstName)
			return err
		}).On(func() error {
			times, err = time.Parse(TimeFormat, createdAt)
			return err
		}).Run()
		if err != nil {
			return nil, err
		}
		task = Task{ID: id, Name: name, CreatedAt: times}
		if status == 0 {
			task.Status = false
		} else {
			task.Status = true
		}
		taskList = append(taskList, task)
	}
	return taskList, nil
}
