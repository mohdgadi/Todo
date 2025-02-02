package main

import "testing"

func TestCreateList(t *testing.T) {
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
	} else {
		t.Logf("list before is %s", list.Name)
		t.Logf("list after is %s", list2.Name)
	}
}
