package main

import "testing"

func TestNegativeCreateList(t *testing.T) {

	list := List{Name: "newlist2"}
	listrepository := SQLiteListRepository{}
	err := listrepository.Create(list)
	if err == nil {
		t.Errorf("List exists but test failed")
		return
	}

}
