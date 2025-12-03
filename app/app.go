package app

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func Start() {
	fmt.Println("Hi, Mum <3")
	//check if json exists, create it if not
	if _, err := os.Stat("tasks.json"); err == nil {
		// path/to/tasks exists
		fmt.Println("The File \"tasks.json\" already exists.")
	} else {
		file, err := os.Create("tasks.json")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("The File \"tasks.json\" has been created.")
		file.Close()
	}

}

func Shutdown() {
	// Shutdown contexts, listeners, and such
}

// store Tasks in a JSON File
func AddTask(id string, description string, status string, createdAt string) {
	//take parameters and write them to the json file
	// set updatedAt string to ""
	taskData := Task{ID: id, Description: description, Status: status, CreatedAt: createdAt, UpdatedAt: ""}

	byteArray, err := json.Marshal(taskData)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))
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

type Task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
