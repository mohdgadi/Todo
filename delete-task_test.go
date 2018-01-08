package main

import "testing"

func TestDeleteExistingTask(t *testing.T) {
	taskID := "12"
	listRepository := SQLiteListRepository{}
	err := listRepository.DeleteTaskFromList(taskID)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
}
func TestDeleteNonExistingTask(t *testing.T) {
	taskID := "32ds"
	listRepository := SQLiteListRepository{}
	err := listRepository.DeleteTaskFromList(taskID)
	if err == nil {
		t.Errorf("Negative test failed")
		return
	}
}
