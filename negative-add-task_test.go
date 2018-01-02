package main

import "testing"

func TestNegativeAddTask(t *testing.T) {

	listrepository := SQLiteListRepository{}
	list := List{Name: "newlist2ew", Tasks: []Task{}}
	task := Task{Name: "taskn"}
	list.Tasks = append(list.Tasks, task)
	err := listrepository.AddTaskToList(list)
	if err == nil {
		t.Errorf("Wrong data but insertion succesful")
	}
}
