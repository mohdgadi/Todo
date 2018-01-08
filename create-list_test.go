package main

import "testing"

func TestCreateNonExistingList(t *testing.T) {
	list := List{Name: "newlist2"}
	listRepository := SQLiteListRepository{}
	err := listRepository.Create(list)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	list2, err := listRepository.Get(list.Name)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if list.Name != list2.Name {
		t.Errorf("Incorrect %s %s", list.Name, list2.Name)
	}
}

func TestCreateExistingList(t *testing.T) {
	list := List{Name: "newlist2"}
	listRepository := SQLiteListRepository{}
	err := listRepository.Create(list)
	if err == nil {
		t.Errorf("List exists but test failed")
		return
	}
}
