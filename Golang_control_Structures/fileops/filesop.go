package filesop

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(Filename string) (float64, error) {
	data, err := os.ReadFile(Filename)

	if err != nil {
		return 1000, errors.New("Failed to find the Balance file.")
	}
	ValueText := string(data)
	Value, err := strconv.ParseFloat(ValueText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored i/p.")
	}
	return Value, nil
}

// Writing the data to the file os package is required
// balance.txt is the file name
//creating the collection of bytes
//0644 is encoding file permission the owner of the file can read and write the file by the other users can only read it

func WriteValueToFile(value float64, Filename string) {
	valuetxt := fmt.Sprint(value)
	os.WriteFile(Filename, []byte(valuetxt), 0644)
}
