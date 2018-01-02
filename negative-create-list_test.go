package main

import "testing"

func TestNegativeCreateList(t *testing.T) {
	list := List{Name: "newlist2"}
	listRepository := SQLiteListRepository{}
	err := listRepository.Create(list)
	if err == nil {
		t.Errorf("List exists but test failed")
		return
	}
}
