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
	for scanner.Scan() {
		cmdLine := scanner.Text()
		cmd := engine.Parse(cmdLine)
		eventLoop.Post(cmd)
	}
	eventLoop.AwaitFinish()
}
