package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var filename string

	fmt.Printf("Enter the name of the text file\n")
	fmt.Scan(&filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error when opening file:\n")
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// read line by line
	for scanner.Scan() {
		cmdLine := scanner.Text()
		sli := strings.Split(cmdLine, " ")
		fmt.Println(sli)
	}
	// handle first encountered error while reading
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error while reading file\n")
		return
	}
}
