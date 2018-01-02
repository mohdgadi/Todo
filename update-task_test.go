package main

import "testing"

func TestUpdateTask(t *testing.T) {
	listName := "List1"
	taskID := "33"
	status := true
	listrepositroy := SQLiteListRepository{}
	err := listrepositroy.UpdateTask(listName, taskID, status)
	if err != nil {
		t.Errorf(err.Error())
	}

}
