package main

import (
	"errors"
	"strconv"
)

// ValidateTaskSpecification used to check if a given task belong to the list.
type ValidateTaskSpecification struct {
	listRepository ListRepository
}

// Enact method checks fetches list by name and checks if given task exists in the List.
func (r ValidateTaskSpecification) Enact(ID string, listName string) error {
	var taskList = []Task{}
	taskList, err := r.listRepository.GetAllTasksFromList(listName)
	if err != nil {
		return err
	}
	for _, task := range taskList {
		if strconv.Itoa(task.ID) == ID {
			return nil
		}
	}
	return errors.New("Doesnt belong")
}
