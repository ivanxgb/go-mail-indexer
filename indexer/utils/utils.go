package utils

import (
	"fmt"
	"os"
)

func ErrorPrinter(err string) {
	fmt.Println("error: " + err)
	os.Exit(1)
}

// FileReader receives a filepath and returns the file content as a string
func FileReader(filepath string) (string, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("There was an error opening the file: " + filepath)
		return "", err
	}

	return string(file), nil
}
