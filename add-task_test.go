package main

import "testing"

func TestAddTask(t *testing.T) {
	listRepository := SQLiteListRepository{}
	list := List{Name: "newlist2", Tasks: []Task{}}
	task := Task{Name: "taskn"}
	list.Tasks = append(list.Tasks, task)
	err := listRepository.AddTaskToList(list)
	if err != nil {
		t.Errorf(err.Error())
	}
}
