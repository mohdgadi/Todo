package main

import "testing"

func TestNegativeGetTask(t *testing.T) {
	listName := "Listxc1"
	ID := "34"
	listRepository := SQLiteListRepository{}
	_, err := listRepository.GetTask(ID, listName)
	if err == nil {
		t.Errorf(err.Error())
		return
	}
}
