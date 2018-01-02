package main

import "testing"

func TestNegativeDeleteAllTask(t *testing.T) {
	listName := "List1"
	listRepository := SQLiteListRepository{}
	err := listRepository.DeleteAllTasksFromList(listName)
	if err == nil {
		t.Errorf("Test failed")
		return
	}
}
