package main

import "testing"

func TestDeleteList(t *testing.T) {
	list := List{Name: "newlist"}
	listrepository := SQLiteListRepository{}
	err := listrepository.Delete(list.Name)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	exists := listrepository.Check(list.Name)
	if exists == true {
		t.Errorf("list exits and not deleted")
	}
}
