package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// To marshal a struct to JSON, its fields must be exported (public) so the encoding/json package can access them.
// The `json` tags provide metadata that tells the JSON encoder how to name the keys in the output JSON.
// Here, the JSON keys will be "title", "content", and "created_at" as specified in the tags.
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Display() {
	fmt.Printf("\nNote title '%v' has the following content: \n\n%v\n\n", note.Title, note.Content)
}

func (note Note) Save() error {
	// Format the title to create a valid filename (e.g., "My Note" → "my_note")
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	// Marshal the note struct to JSON
	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	// Write the JSON data to a file named after the note title
	// The file will have 0644 permissions (rw-r--r--)
	return os.WriteFile(fileName, json, 0644)
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input") // Returning dummy note in case of validation error
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
