package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func Start() {
	// fmt.Println("Hi, Mum <3")
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

		openFile, err := os.OpenFile("tasks.json", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		if _, err := openFile.Write([]byte("[]")); err != nil {
			log.Fatal(err)
		}
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

	//create json object to write in file
	byteArray, err := json.Marshal(taskData)
	if err != nil {
		fmt.Println(err)
	}

	file, err := os.OpenFile("tasks.json", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	//write inside the array
	content, err := os.ReadFile("tasks.json")
	if err != nil {
		panic(err)
	}
	//find the [ char
	pos := bytes.IndexByte(content, '[')
	if pos == -1 {
		println("[ nicht gefunden")
		return
	}

	newContent := make([]byte, 0, len(content)+len(byteArray))

	newContent = append(newContent, content[:pos+1]...) // file data until [
	newContent = append(newContent, byteArray...)       // new task object

	//check if comma needed
	if len(content) > 3 {
		newContent = append(newContent, []byte(",")...)
		fmt.Println("comma inserted")
	}
	newContent = append(newContent, content[pos+1:]...) // rest of the file

	err = os.WriteFile("tasks.json", newContent, 0644)
	if err != nil {
		panic(err)
	}
}

// retrieve the task that one wants to update, change values except createdAt
func UpdatedTask(id string, description string, status string, createdAt string, updatedAt string) {

	//taskData := Task{ID: id, Description: description, Status: status, CreatedAt: createdAt, UpdatedAt: ""}
	taskData := Task{ID: id, Description: description, Status: status, CreatedAt: createdAt, UpdatedAt: updatedAt}
	byteArray, err := json.Marshal(taskData)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))
}

// delete task from json by id
func DeleteTask(id string) {

	//open json file
	//find taskobjekt with id
	//delete entry
	//save file
}

// change status of task object to "in-progress" by id
func markInProgress(id string) {

}

// change status of task object to "done" by id
func markDone(id string) {

}

// show all tasks from the json in list
func listAll() {

}

// show all tasks with status "done" from the json in list
func listDone() {

}

// show all tasks with status "todo" from the json in list
func listToDo() {

}

// show all tasks with status "in-progress" from the json in list
func listInProgress() {

}

type Task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
