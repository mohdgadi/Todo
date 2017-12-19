package main

type ListModel struct {
	ID        int
	listName  string `gorm:"primary_key"`
	createdAt string
}
