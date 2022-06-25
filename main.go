package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Command interface {
	Execute()
}

type printCommand string

func (pc printCommand) Execute() {
	fmt.Println(pc)
}

func parse(cmdLine string) (Command, error) {
	sli := strings.Split(cmdLine, " ")
	if sli[0] == "print" {
		return printCommand(strings.Join(sli[1:], " ")), nil
	} else {
		return nil, fmt.Errorf("ERROR: This command doesn't exist.")
	}
}

func main() {
	var commands []Command
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
		cmd, errParse := parse(cmdLine)
		if errParse != nil {
			fmt.Printf("%s Row: %d\n", errParse.Error(), row)
			continue
		}
		commands = append(commands, cmd)
		row++
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
