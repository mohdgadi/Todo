package main

import "testing"

func TestNegativeDeleteList(t *testing.T) {
	list := List{Name: "newlistsxxs"}
	listRepository := SQLiteListRepository{}
	err := listRepository.Delete(list.Name)
	if err == nil {
		t.Errorf("test failed as list deleted succesful")
		return
	}
}
