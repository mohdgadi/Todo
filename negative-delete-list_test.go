package main

import "testing"

func TestNegativeDeleteList(t *testing.T) {
	list := List{Name: "newlistsxxs"}
	listrepository := SQLiteListRepository{}
	err := listrepository.Delete(list.Name)
	if err == nil {
		t.Errorf("test failed as list deleted succesful")
		return
	}

}
