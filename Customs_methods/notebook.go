// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"

// 	"example.com/notebook/note"
// 	"example.com/notebook/todo"
// )

// // INTEFACE : is like a contract that defines a set of actions(methods) that a type must be able to peform all the tasks listed
// // In Interface we are only defining the methods and we are not giving any body of the method (eg.: below we are only metioning the method "save" and telling the return type and parameters , here we are not passing any parameters because the save method is not taking any )

// // Defining the Interface or creating the Interface
// type saver interface {
// 	Save() error
// }

// ///CREATING AN EMBEDDED INTERFACE

// type outputtable interface {
// 	saver // embedding the saver
// 	Display()
// }

// func main() {
// 	title, content := getNoteData()
// 	// title and content variables are calling the getNotedata

// 	todoText := getUserInput("Todo text : ")

// 	// here the todoText is reciving the text from the "gteUserinput" function

// 	todo, err := todo.New(todoText) // here we are using the "todo" is calling the consturctor function named "new" from the "todo" package

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	userNote, err := note.New(title, content)

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	err = outputData(todo)

// 	if err != nil {
// 		return
// 	}

// 	err = outputData(userNote)

// }

// func printSomething(value interface{}) {
// 	fmt.Println(value)
// }

// func outputData(data outputtable) error {
// 	data.Display()
// 	return saveData(data)
// }

// func saveData(data saver) error {
// 	err := data.Save()

// 	if err != nil {
// 		fmt.Println("Saving the note failed.")
// 		return err
// 	}

// 	fmt.Println("Saving the note succeeded!")
// 	return nil
// }

// func getTodoData() string {
// 	return getUserInput("Todo text : ")
// }

// func getNoteData() (string, string) {
// 	title := getUserInput("Note title:")
// 	content := getUserInput("Note content:")

// 	return title, content
// }

// func getUserInput(prompt string) string {
// 	fmt.Printf("%v ", prompt)

// 	reader := bufio.NewReader(os.Stdin)

// 	text, err := reader.ReadString('\n')

// 	if err != nil {
// 		return ""
// 	}

// 	text = strings.TrimSuffix(text, "\n")
// 	text = strings.TrimSuffix(text, "\r")

// 	return text
// }
