package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Define the file name
	fileName := "example.txt"

	// Read the content of the file
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Display the file content
	fmt.Println("File Content:")
	fmt.Println(string(content))

	// Append new content to the file
	newContent := "\nThis is some new content appended to the file."
	err = appendToFile(fileName, newContent)
	if err != nil {
		fmt.Printf("Error appending to file: %v\n", err)
		return
	}

	fmt.Println("New content has been appended to the file.")
}

func appendToFile(fileName, content string) error {
	// Open the file in append mode with write permissions
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the content to the file
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
