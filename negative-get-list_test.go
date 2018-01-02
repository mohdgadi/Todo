package main

import "testing"

func TestNegativeGetList(t *testing.T) {
	listName := "List1xs"
	listRepository := SQLiteListRepository{}
	_, err := listRepository.Get(listName)
	if err == nil {
		t.Errorf("List doesnt exist but returned")
		return
	}
}
