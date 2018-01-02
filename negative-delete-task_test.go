package main

import "testing"

func TestNegativeDeleteTask(t *testing.T) {
	listName := "List1"
	taskID := "32ds"
	listRepository := SQLiteListRepository{}
	err := listRepository.DeleteTask(taskID, listName)
	if err == nil {
		t.Errorf("Negative test failed")
		return
	}
}
