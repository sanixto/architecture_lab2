package main

import (
	"bufio"
	"eventloop/engine"
	"fmt"
	"os"
)

func main() {
	var filename string

	fmt.Printf("Enter the name of the text file\n")
	fmt.Scan(&filename)

	eventLoop := new(engine.Loop)
	eventLoop.Start()

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
		cmd := engine.Parse(cmdLine, row)
		row++
		eventLoop.Post(cmd)
	}
	// handle first encountered error while reading
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error while reading file\n")
		return
	}
	eventLoop.AwaitFinish()
}
