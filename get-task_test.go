package main

import "testing"

func TestGetExistingTask(t *testing.T) {
	ID := "34"
	listRepository := SQLiteListRepository{}
	task, err := listRepository.GetTaskFromList(ID)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if task.ID != 34 {
		t.Errorf("Get task failed")
	}
}
func TestGetNonExistingTask(t *testing.T) {
	ID := "34"
	listRepository := SQLiteListRepository{}
	_, err := listRepository.GetTaskFromList(ID)
	if err == nil {
		t.Errorf(err.Error())
		return
	}
}
