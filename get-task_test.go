package main

import "testing"

func TestGetTask(t *testing.T) {
	listName := "List1"
	ID := "34"
	listRepository := SQLiteListRepository{}
	task, err := listRepository.GetTask(ID, listName)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if task.ID != 34 {
		t.Errorf("Get task failed")
	}

}
