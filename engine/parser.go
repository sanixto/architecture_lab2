package engine

import (
	"fmt"
	"strconv"
	"strings"
)

func Parse(cmdLine string) (Command, error) {
	sli := strings.Split(cmdLine, " ")
	length := len(sli)

	switch length {
	case 0:
		return nil, fmt.Errorf("ERROR: The empty string.")
	case 1:
		if sli[0] == "print" || sli[0] == "printc" {
			return nil, fmt.Errorf("ERROR: No arguments.")
		}
		return nil, fmt.Errorf("ERROR: This command doesn't exist.")
	case 2:
		if sli[0] == "print" {
			return printCommand(sli[1]), nil
		}
		if sli[0] == "printc" {
			return nil, fmt.Errorf("ERROR: Too few arguments.")
		}
		return nil, fmt.Errorf("ERROR: This command doesn't exist.")
	case 3:
		if sli[0] == "print" {
			return nil, fmt.Errorf("ERROR: Too many arguments.")
		}
		if sli[0] == "printc" {
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
	default:
		return nil, fmt.Errorf("ERROR: Too many arguments.")
	}
}
