package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

///---------This whole Block of code is the constrution of the struct------//

type Todo struct {
	Text string `json:"content"`
}

///---------This whole Block of code Below is adding the "Method"------//

func (todo Todo) Display() {

	fmt.Println(todo.Text)
	/// --- Here the "todo" is the name of the struct and we are accesing the variables of that struct by the "." operator
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

// ----------- The below block of code we are creating an instance to the struct ---------//

func New(content string) (Todo, error) {

	if content == "" {
		return Todo{}, errors.New("Invalid Input")
	}

	return Todo{

		Text: content,
	}, nil

}
