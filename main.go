package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router := mux.NewRouter()
	taskRepository := SQLiteTaskRepository{}
	listRepository := SQLiteListRepository{taskRepository}
	getListIntent := GetListIntent{listRepository}
	createListIntent := CreateListIntent{listRepository}
	deleteListIntent := DeleteListIntent{listRepository}
	deleteTaskIntent := DeleteTaskIntent{taskRepository}
	updateTaskIntent := UpdateTaskIntent{taskRepository}
	getTaskIntent := GetTaskIntent{taskRepository}
	addTaskIntent := AddTaskIntent{taskRepository, listRepository}

	router.HandleFunc("/lists", createListIntent.Enact).Methods("POST")
	router.HandleFunc("/lists/{id}", deleteListIntent.Enact).Methods("DELETE")
	router.HandleFunc("/lists/{id}", getListIntent.Enact).Methods("GET")
	router.HandleFunc("/tasks/{id}", getTaskIntent.Enact).Methods("GET")
	router.HandleFunc("/tasks", addTaskIntent.Enact).Methods("POST")
	router.HandleFunc("/tasks/{id}", deleteTaskIntent.Enact).Methods("DELETE")
	router.HandleFunc("/tasks", updateTaskIntent.Enact).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// Use decoder:done
//Add comments:done
//Use camel casing:done
//ReName getall and other funcctions:done
//check rows.next:done
//http write in one line:done
//return error in one line:done
//break line:done
//define variables in a single block:done
//gitignore.io:done
//enviroment constants:done
//change dependencies:done
