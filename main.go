package main

import (
	"bufio"
	"eventloop/engine"
	"fmt"
	"os"
)

func main() {
	var commands []engine.Command
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
	var row int
	for scanner.Scan() {
		cmdLine := scanner.Text()
		cmd, errParse := engine.Parse(cmdLine)
		row++
		if errParse != nil {
			fmt.Printf("%s Row: %d\n", errParse.Error(), row)
			continue
		}
		commands = append(commands, cmd)
	}
	// handle first encountered error while reading
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error while reading file\n")
		return
	}

	for _, v := range commands {
		v.Execute()
	}
}
