package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Command interface {
	Execute()
}

type printCommand string

func (p printCommand) Execute() {
	fmt.Println(p)
}

type printcCommand struct {
	count  int
	symbol string
}

func (p *printcCommand) Execute() {
	for i := 0; i < p.count; i++ {
		fmt.Print(p.symbol)
	}
	fmt.Println()
}

func parse(cmdLine string) (Command, error) {
	sli := strings.Split(cmdLine, " ")
	if sli[0] == "print" {
		return printCommand(strings.Join(sli[1:], " ")), nil
	}
	if sli[0] == "printc" {
		if len(sli) > 3 {
			return nil, fmt.Errorf("ERROR: Too many arguments.")
		}
		count, errC := strconv.Atoi(sli[1])
		if errC != nil {
			return nil, fmt.Errorf("ERROR: The second argument has to be a number.")
		}
		symbol := sli[2]
		if len(symbol) > 1 {
			return nil, fmt.Errorf("ERROR: The third argument has to be a symbol.")
		}
		return &printcCommand{count: count, symbol: symbol}, nil
	}
	return nil, fmt.Errorf("ERROR: This command doesn't exist.")
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
