package main

import (
	"database/sql"
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

// TaskRep ...
type TaskRep interface {
	add(task Task, listname string) (int, error)
	delete(ID string) (int, error)
	deletetask(listname string) error
	get(ID string) (Task, error)
	update(task Task) (int, error)
}

//TaskRepo ...
type TaskRepo struct {
}

func (t TaskRepo) add(tasks Task, listname string) (int, error) {
	listrepo := ListRepo{}
	if listrepo.check(listname) == false {
		fmt.Println("List doesnt exist")
		return 0, nil
	}
	fmt.Println("reached")
	task := TaskModel{name: tasks.Name, status: false, listName: listname}
	task.createdAt = time.Now().Local().Format("2006-01-02")

	database, _ := sql.Open("sqlite3", "./test.db")
	defer database.Close()

	var query = "INSERT INTO tasks (createdat,name,status,listname) VALUES ( '" + task.createdAt + "','" + task.name + "','0','" + task.listName + "')"
	statement, err := database.Prepare(query)
	statement.Exec()
	if err != nil {
		return 0, err

	}
	return 1, nil

}

func (t TaskRepo) delete(ID string) (int, error) {
	database, _ := sql.Open("sqlite3", "./test.db")
	defer database.Close()

	var query = "DELETE FROM tasks WHERE ID = '" + ID + "' ;"
	statement, err := database.Prepare(query)
	res, err := statement.Exec()
	affected, err := res.RowsAffected()
	if affected == 0 || err != nil {
		return 0, err
	}
	return 1, nil
}
func (t TaskRepo) get(ID string) (Task, error) {
	database, _ := sql.Open("sqlite3", "./test.db")
	defer database.Close()
	var query = "SELECT * FROM tasks WHERE ID = '" + ID + "' ;"
	rows, err := database.Query(query)
	var task Task
	if err != nil {
		return task, err
	}

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

	}
	return task, nil
}

func (t TaskRepo) deletelist(listname string) (int, error) {
	database, _ := sql.Open("sqlite3", "./test.db")
	defer database.Close()
	var query = "DELETE FROM tasks WHERE listname = '" + listname + "';"
	statement, err := database.Prepare(query)
	statement.Exec()
	if err != nil {
		return 0, err
	}
	fmt.Println("deleted task list")
	return 0, nil
}

func (t TaskRepo) update(task Task) (int, error) {
	database, _ := sql.Open("sqlite3", "./test.db")
	defer database.Close()
	var status string
	if task.Status == false {
		status = "0"
	} else {
		status = "1"
	}
	var query = "UPDATE tasks SET status=" + status + "WHERE ID=" + strconv.Itoa(task.ID) + ";"
	statement, err := database.Prepare(query)
	res, err := statement.Exec()
	affected, err := res.RowsAffected()
	if affected == 0 || err != nil {
		return 0, err
	}
	return 1, nil
}
