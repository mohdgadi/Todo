package main

import "testing"

func TestDeleteList(t *testing.T) {
	list := List{Name: "newlist"}
	listRepository := SQLiteListRepository{}
	err := listRepository.Delete(list.Name)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	exists := listRepository.Check(list.Name)
	if exists == true {
		t.Errorf("list exits and not deleted")
	}
}
