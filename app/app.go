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
func AddTask(id int, description string, status string, createdAt string) error {
	//take parameters and write them to the json file
	// set updatedAt string to ""
	taskData := Task{ID: id, Description: description, Status: status, CreatedAt: createdAt, UpdatedAt: ""}

	//create json object to write in file
	byteArray, err := json.Marshal(taskData)
	if err != nil {
		return err
	}

	insertAfterNthOccurrenceStream("tasks.json", '[', byteArray, 1)
	return nil
}

// retrieve the task that one wants to update, change values except createdAt
func UpdateTask(id int, description string, status string, updatedAt string) error {

	//taskData := Task{ID: id, Description: description, Status: status, CreatedAt: createdAt, UpdatedAt: ""}
	taskData := Task{ID: id, Description: description, Status: status, UpdatedAt: updatedAt}

	// byteArray, err := json.Marshal(taskData)
	// if err != nil {
	// 	return err
	// }

	ReplaceTypedJSONObjectByID("tasks.json", id, taskData)
	return nil
}

// delete task from json by id
func DeleteTask(id int) {
	DeleteTypedJSONObjectByID("tasks.json", id)

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

func ReplaceTypedJSONObjectByID(filePath string, id int, newObject Task) error {
	// JSON-Datei einlesen
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("fehler beim Lesen der Datei: %w", err)
	}

	// JSON in Array von Items parsen
	var items []Task
	if err := json.Unmarshal(data, &items); err != nil {
		return fmt.Errorf("fehler beim Parsen des JSON: %w", err)
	}

	// Nach Item mit der ID suchen und ersetzen
	found := false
	for i, item := range items {
		if item.ID == id {
			// createdAt vom alten Item übernehmen (falls vorhanden)
			if item.CreatedAt != "" {
				newObject.CreatedAt = item.CreatedAt
			}
			items[i] = newObject
			found = true

			break
		}
	}

	if !found {
		return fmt.Errorf("objekt mit ID %d nicht gefunden", id)
	}

	// JSON zurück in Datei schreiben
	newData, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return fmt.Errorf("fehler beim Erstellen des JSON: %w", err)
	}

	if err := os.WriteFile(filePath, newData, 0644); err != nil {
		return fmt.Errorf("fehler beim Schreiben der Datei: %w", err)
	}

	return nil
}

func DeleteTypedJSONObjectByID(filePath string, id int) error {
	// JSON-Datei einlesen
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("fehler beim Lesen der Datei: %w", err)
	}

	// JSON in Array von Items parsen
	var items []Task
	if err := json.Unmarshal(data, &items); err != nil {
		return fmt.Errorf("fehler beim Parsen des JSON: %w", err)
	}

	// Nach Item mit der ID suchen und ersetzen
	found := false
	for i, item := range items {
		if item.ID == id {

			items[i] = items[len(items)-1]

			found = true

			break
		}
	}

	filteredItems := items[:len(items)-1]

	if !found {
		return fmt.Errorf("objekt mit ID %d nicht gefunden", id)
	}

	// JSON zurück in Datei schreiben
	newData, err := json.MarshalIndent(filteredItems, "", "  ")
	if err != nil {
		return fmt.Errorf("fehler beim Erstellen des JSON: %w", err)
	}

	if err := os.WriteFile(filePath, newData, 0644); err != nil {
		return fmt.Errorf("fehler beim Schreiben der Datei: %w", err)
	}

	return nil
}

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
