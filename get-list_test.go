package main

import "testing"

func TestGetList(t *testing.T) {
	listName := "List1"
	listrepository := SQLiteListRepository{}
	_, err := listrepository.Get(listName)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

}
