package main

import "testing"

func TestAddTask(t *testing.T) {

	listrepository := SQLiteListRepository{}
	list := List{Name: "newlist2", Tasks: []Task{}}
	task := Task{Name: "taskn"}
	list.Tasks = append(list.Tasks, task)
	err := listrepository.AddTaskToList(list)
	if err != nil {
		t.Errorf(err.Error())
	}
}
