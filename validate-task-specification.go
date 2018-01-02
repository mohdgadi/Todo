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
	list, err := r.listRepository.Get(listName)
	if err != nil {
		return err
	}
	tasks := list.Tasks
	for _, task := range tasks {
		if strconv.Itoa(task.ID) == ID {
			return nil
		}
	}
	return errors.New("Doesnt belong")
}
