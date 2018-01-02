package main

import "testing"

func TestGetTask(t *testing.T) {
	listName := "List1"
	ID := "34"
	listrepository := SQLiteListRepository{}
	task, err := listrepository.GetTask(ID, listName)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if task.ID != 34 {
		t.Errorf("Get task failed")
	}

}
