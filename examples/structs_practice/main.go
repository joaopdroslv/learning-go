package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}

	fmt.Println("Saving the note succeeded!")
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
