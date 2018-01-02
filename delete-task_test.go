package main

import "testing"

func TestDeleteTask(t *testing.T) {
	listName := "List1"
	taskID := "32"
	listrepository := SQLiteListRepository{}
	err := listrepository.DeleteTask(taskID, listName)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	tasks, err := listrepository.GetTask(taskID, listName)
	if err != nil {
		t.Errorf(err.Error())
	} else {
		if tasks.ID != 0 {
			t.Errorf("Task deletion failed")
		}
	}
}
