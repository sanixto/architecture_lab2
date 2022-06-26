package engine

import (
	"strconv"
	"strings"
)

func Parse(cmdLine string) Command {
	if len(cmdLine) == 0 {
		return printCommand("ERROR: The empty string.")
	}
	sli := strings.Fields(cmdLine)
	length := len(sli)

	switch length {
	case 0:
		return printCommand("ERROR: The empty string.")
	case 1:
		if sli[0] == "print" || sli[0] == "printc" {
			return printCommand("ERROR: No arguments.")
		}
		return printCommand("ERROR: This command doesn't exist.")
	case 2:
		if sli[0] == "print" {
			return printCommand(sli[1])
		}
		if sli[0] == "printc" {
			return printCommand("ERROR: Too few arguments.")
		}
		return printCommand("ERROR: This command doesn't exist.")
	case 3:
		if sli[0] == "print" {
			return printCommand("ERROR: Too many arguments.")
		}
		if sli[0] == "printc" {
			count, errC := strconv.Atoi(sli[1])
			if errC != nil {
				return printCommand("ERROR: The second argument has to be a number.")
			}
			symbol := sli[2]
			if len(symbol) > 1 {
				return printCommand("ERROR: The third argument has to be a symbol.")
			}
			return printcCommand{count: count, symbol: symbol}
		}
		return printCommand("ERROR: This command doesn't exist.")
	default:
		return printCommand("ERROR: Too many arguments.")
	}
}
