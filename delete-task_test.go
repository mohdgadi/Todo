package main

import "testing"

func TestDeleteExistingTask(t *testing.T) {
	taskID := "32"
	listRepository := SQLiteListRepository{}
	err := listRepository.DeleteTaskFromList(taskID)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	tasks, err := listRepository.GetTaskFromList(taskID)
	if err != nil {
		t.Errorf(err.Error())
	} else {
		if tasks.ID != 0 {
			t.Errorf("Task deletion failed")
		}
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
