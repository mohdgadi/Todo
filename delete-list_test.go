package main

import "testing"

func TestDeleteExistingList(t *testing.T) {
	list := List{Name: "newlist"}
	listRepository := SQLiteListRepository{}
	err := listRepository.Delete(list.Name)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	exists := listRepository.CheckIfExists(list.Name)
	if exists == true {
		t.Errorf("list exits and not deleted")
	}
}
func TestDeleteNonExistingList(t *testing.T) {
	list := List{Name: "newlistsxxs"}
	listRepository := SQLiteListRepository{}
	err := listRepository.Delete(list.Name)
	if err == nil {
		t.Errorf("test failed as list deleted succesful")
		return
	}
}
