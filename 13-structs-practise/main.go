package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

type outputtable interface {
	saver 
	Display()
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("content: ")
	return title, content
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)
	
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// todo.Display()
	// err = saveData(todo)
	err = outputData(todo)
	if err != nil {
		return 
	}

	// err = todo.Save()
	// if err != nil {
	// 	fmt.Println("Saving the failed failed.")
	// 	return
	// }
  // fmt.Println("Saving the todo succeded")

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// userNote.Display()
	// err = saveData(userNote)
	err = outputData(userNote)
	if err != nil {
		return
	}

	// err = userNote.Save()
	// if err != nil {
	// 	fmt.Println("Saving the note failed.")
	// 	return
	// }
  // fmt.Println("Saving the note succeded")


}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

  fmt.Println("Saving the note succeded")
	return nil
}

func getUserInput(prompt string) string {
	fmt.Printf("%v", prompt)
	// var value string
	// fmt.Scanln(&value)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
