package app

import (
	"fmt"
)

func Start() {
	fmt.Println("Hi, Mum <3")

}

func Shutdown() {
	// Shutdown contexts, listeners, and such
}

// store Tasks in a JSON File
func addTask(id string, description string, status string, createdAt string) {
	//take parameters and write them to the json file
	// set updatedAt string to ""

}

func updatedTask(id string, description string, status string, createdAt string, updatedAt string) {
	// retrieve the task that one wants to update, change values except createdAt
}

func deleteTask(id string) {
	//delete task from json by id
}

func markInProgress(id string) {

}

func markDone(id string) {

}

func listAll() {
	//show all tasks in list
}

func listDone() {

}

func listToDo() {

}

func listInProgress() {

}
