package main

import "testing"

func TestNegativeGetAllTaskFromList(t *testing.T) {
	listName := "List1"
	listRepository := SQLiteListRepository{}
	_, err := listRepository.GetAllTasksFromList(listName)
	if err == nil {
		t.Errorf("Test failed")
		return
	}
}
