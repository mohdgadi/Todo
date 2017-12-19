package main

type TaskModel struct {
	ID        int `gorm:"primary_key"`
	name      string
	status    bool
	listModel ListModel
	listName  string
	createdAt string
}
