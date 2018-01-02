package main

import "testing"

func TestDeleteAllTask(t *testing.T) {
	listName := "Lisdt1"
	listRepository := SQLiteListRepository{}
	err := listRepository.DeleteAllTasksFromList(listName)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
}
