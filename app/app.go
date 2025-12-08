package app

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
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

	insertAfterNthOccurrenceStream("tasks.json", '[', byteArray, 1)
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

func insertAfterNthOccurrenceStream(filePath string, search byte, textToInsert []byte, n int) error {
	// Eingabedatei öffnen
	inputFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	// Temporäre Datei im gleichen Verzeichnis erstellen
	tempFile, err := os.CreateTemp("", "insert_temp_*.txt")
	if err != nil {
		return err
	}
	tempPath := tempFile.Name()
	defer os.Remove(tempPath) // Temporäre Datei am Ende löschen

	reader := bufio.NewReader(inputFile)
	writer := bufio.NewWriter(tempFile)

	count := 0
	inserted := false

	// Lesen und in temporäre Datei schreiben
	for {
		b, err := reader.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			tempFile.Close()
			return err
		}

		// Byte in temporäre Datei schreiben
		if err := writer.WriteByte(b); err != nil {
			tempFile.Close()
			return err
		}

		// Prüfen ob es das gesuchte Zeichen ist
		if b == search && !inserted {
			count++
			if count == n {
				// Nächstes Zeichen anschauen um zu prüfen ob schon Inhalt da ist
				nextByte, err := reader.ReadByte()
				if err != nil && err != io.EOF {
					tempFile.Close()
					return err
				}

				// Wenn das nächste Zeichen kein ] ist, dann ist schon Inhalt vorhanden
				if err != io.EOF && nextByte != ']' {
					// Komma vor dem neuen Text einfügen
					if _, err := writer.Write(textToInsert); err != nil {
						tempFile.Close()
						return err
					}
					if err := writer.WriteByte(','); err != nil {
						tempFile.Close()
						return err
					}
				} else {
					// Kein Inhalt vorhanden, nur neuen Text einfügen
					if _, err := writer.Write(textToInsert); err != nil {
						tempFile.Close()
						return err
					}
				}

				// Das bereits gelesene nächste Byte auch schreiben
				if err != io.EOF {
					if err := writer.WriteByte(nextByte); err != nil {
						tempFile.Close()
						return err
					}
				}

				inserted = true
			}
		}
	}

	// Puffer leeren und Dateien schließen
	if err := writer.Flush(); err != nil {
		tempFile.Close()
		return err
	}
	tempFile.Close()
	inputFile.Close()

	// Temporäre Datei über Originaldatei kopieren
	if err := os.Rename(tempPath, filePath); err != nil {
		return err
	}

	return nil
}

type Task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
