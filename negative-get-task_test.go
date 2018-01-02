package main

import "testing"

func TestNegativeGetTask(t *testing.T) {
	listName := "Listxc1"
	ID := "34"
	listrepository := SQLiteListRepository{}
	_, err := listrepository.GetTask(ID, listName)
	if err == nil {
		t.Errorf(err.Error())
		return
	}

}
