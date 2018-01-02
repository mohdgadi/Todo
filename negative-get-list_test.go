package main

import "testing"

func TestNegativeGetList(t *testing.T) {
	listName := "List1xs"
	listrepository := SQLiteListRepository{}
	_, err := listrepository.Get(listName)
	if err == nil {
		t.Errorf("List doesnt exist but returned")
		return
	}
}
