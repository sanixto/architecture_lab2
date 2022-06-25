package main

import (
	"bufio"
	"eventloop/engine"
	"fmt"
	"os"
	"strconv"
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
		cmd, errParse := engine.Parse(cmdLine)
		row++
		if errParse != nil {
			str := "print " + errParse.Error() + " Row: " + strconv.Itoa(row)
			cmd, _ = engine.Parse(str)
		}
		eventLoop.Post(cmd)
	}
	// handle first encountered error while reading
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error while reading file\n")
		return
	}
	eventLoop.AwaitFinish()
}
