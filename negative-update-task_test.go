package main

import "testing"

func TestNegativeUpdateTask(t *testing.T) {
	listName := "List1"
	taskID := "34s"
	status := true
	listRepositroy := SQLiteListRepository{}
	err := listRepositroy.UpdateTask(listName, taskID, status)
	if err == nil {
		t.Errorf("Test failed for updation")
	}
}
