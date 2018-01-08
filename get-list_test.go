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
func TestNegativeGetList(t *testing.T) {
	listName := "List1xs"
	listRepository := SQLiteListRepository{}
	_, err := listRepository.Get(listName)
	if err == nil {
		t.Errorf("List doesnt exist but returned")
		return
	}
}
