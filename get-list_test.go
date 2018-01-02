package main

import "testing"

func TestGetList(t *testing.T) {
	listName := "List1"
	listRepository := SQLiteListRepository{}
	_, err := listRepository.Get(listName)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

}
