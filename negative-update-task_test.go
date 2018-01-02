package main

import "testing"

func TestNegativeUpdateTask(t *testing.T) {
	listName := "List1"
	taskID := "34s"
	status := true
	listrepositroy := SQLiteListRepository{}
	err := listrepositroy.UpdateTask(listName, taskID, status)
	if err == nil {
		t.Errorf("Test failed for updation")
	}

}
