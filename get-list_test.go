package main

import "testing"

func TestGetList(t *testing.T) {
	listName := "List1"
	listRepository := SQLiteListRepository{}
	list, err := listRepository.Get(listName)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if list.Name == "" {
		t.Errorf(err.Error())
		return
	}
}
