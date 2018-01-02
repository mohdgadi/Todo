package main

import "testing"

func TestNegativeAddTask(t *testing.T) {
	listRepository := SQLiteListRepository{}
	list := List{Name: "newlist2ew", Tasks: []Task{}}
	task := Task{Name: "taskn"}
	list.Tasks = append(list.Tasks, task)
	err := listRepository.AddTaskToList(list)
	if err == nil {
		t.Errorf("Wrong data but insertion succesful")
	}
}
