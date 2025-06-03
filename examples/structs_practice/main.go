package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
	"example.com/notes/todo"
)

// The 'saver' interface defines a contract that requires implementing types
// to have a Save method, which returns an error. This ensures consistency
// across different entities that need saving behavior.
type saver interface {
	Save() error // Method signature that must be implemented by conforming types
}

// type displayer interface {
// 	Display()
// }

// type outputtable interface {
// 	Save() error
// 	Display()
// }

// Embedded interface:
// Any type that implements outputtable must implement all methods from the saver interface
// as well as the Display() method.
type outputtable interface {
	saver     // Embeds the saver interface
	Display() // Additional method required by outputtable
}

func main() {
	printAnything(1)
	printAnything(1.5)
	printAnything("Something!")

	title, content := getNoteData()
	todoText := getUserInput("Todo text:")

	userTodo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userTodo)

	if err != nil {
		return
	}

	_ = outputData(userNote) // Ignore this return, the application will end anyways
}

// The empty interface (interface{}) allows this function to accept any type,
// since all types in Go implement at least the empty interface.
func printAnything(value interface{}) {
	// After validation, if it'a int, Go automatically types the typedValue as an int
	intVal, ok := value.(int) // Syntax to validate a value type
	if ok {
		// Now we can perform operations with the now correct typed value
		_ = intVal + 1 // For example
		fmt.Println("It's now typed as an int!")
	}

	floatVal, ok := value.(float64)
	if ok {
		_ = floatVal + 0.99
		fmt.Println("It's now typed as a float64!")
	}

	stringVal, ok := value.(string)
	if ok {
		fmt.Println("It's now typed as a string!", stringVal)
	}

	// Using switch
	// switch value.(type) { // Syntax to validate type VARIABLE.(type) inside a switch statement
	// case int:
	// 	fmt.Println("It's a Integer! |", value)
	// case float64:
	// 	fmt.Println("It's a float64! |", value)
	// case string:
	// 	fmt.Println("It's a string! |", value)
	// }
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

// By accepting a 'saver' interface, we ensure that any type passed to this function
// has a Save method. This promotes flexibility and enforces consistent behavior.
func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	fmt.Println("Saving succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	// Use bufio.Reader for reading multi-word or full-line input.
	// fmt.Scan() stops at spaces, so it's not suitable for full lines.
	reader := bufio.NewReader(os.Stdin)

	// Read until newline character ('\n') is encountered
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	// Remove newline characters for cross-platform compatibility (\n on Unix, \r\n on Windows)
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
