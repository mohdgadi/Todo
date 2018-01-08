package main

import "testing"

func TestUpdateExistingTask(t *testing.T) {
	var tasklist []Task
	list := List{Name: "List1"}
	task := Task{ID: 33, Status: true}
	tasklist = append(tasklist, task)
	listrepositroy := SQLiteListRepository{}
	err := listrepositroy.UpdateTaskStatus(list)
	if err != nil {
		t.Errorf(err.Error())
	}
}
func TestUpdateNonExistingTask(t *testing.T) {
	var tasklist []Task
	list := List{Name: "List1"}
	task := Task{ID: 300, Status: true}
	tasklist = append(tasklist, task)
	listRepositroy := SQLiteListRepository{}
	err := listRepositroy.UpdateTaskStatus(list)
	if err == nil {
		t.Errorf("Test failed for updation")
	}
}
