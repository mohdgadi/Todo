package main

import "testing"

func TestGetAllTaskFromList(t *testing.T) {
	listName := "List1"
	listRepository := SQLiteListRepository{}
	_, err := listRepository.GetAllTasksFromList(listName)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
}
