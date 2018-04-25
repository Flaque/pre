package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func append(content string, filename string) string {

	// Read
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file '%s'\n", err.Error())
		os.Exit(1)
	}

	// Output
	return content + string(b)
}

const helpText = "'pre' listens to stdin and appends it to a file.\n\n" +
	"Usage: pre FILENAME \n" +
	"Example: 'cat puppers.txt | pre cute-things.txt'"

func main() {
	if len(os.Args) < 2 {
		fmt.Println(helpText)
		os.Exit(1)
	}

	// Check if file exists
	if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
		fmt.Printf("File '%s' does not exist.\n", os.Args[1])
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)

	total := ""
	for scanner.Scan() {
		total += scanner.Text() + "\n"
	}
	total += "\n"

	output := append(total, os.Args[1])
	fmt.Println(output)
}
