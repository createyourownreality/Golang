package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

///---------This whole Block of code is the constrution of the struct------//

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

///---------This whole Block of code Below is adding the "Method"------//

func (note Note) Display() {

	fmt.Printf("Your note titled %v has the follwing content: \n\n%v\n\n", note.Title, note.Content)
	/// --- Here the "note" is the name of the struct and we are accesing the variables of that struct by the "." operator
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, "", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

// ----------- The below block of code we are creating an instance to the struct ---------//

func New(title, content string) (Note, error) {

	if title == "" || content == "" {
		return Note{}, errors.New("Invalid Input")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil

}
