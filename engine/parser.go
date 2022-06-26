package engine

import (
	"strconv"
	"strings"
)

func Parse(cmdLine string, row int) Command {
	if len(cmdLine) == 0 {
		return printCommand("ERROR: The empty string. Row: " + strconv.Itoa(row))
	}
	sli := strings.Split(cmdLine, " ")
	length := len(sli)

	switch length {
	case 0:
		return printCommand("ERROR: The empty string. Row: " + strconv.Itoa(row))
	case 1:
		if sli[0] == "print" || sli[0] == "printc" {
			return printCommand("ERROR: No arguments.")
		}
		return printCommand("ERROR: This command doesn't exist. Row: " + strconv.Itoa(row))
	case 2:
		if sli[0] == "print" {
			return printCommand(sli[1])
		}
		if sli[0] == "printc" {
			return printCommand("ERROR: Too few arguments. Row: " + strconv.Itoa(row))
		}
		return printCommand("ERROR: This command doesn't exist. Row: " + strconv.Itoa(row))
	case 3:
		if sli[0] == "print" {
			return printCommand("ERROR: Too many arguments. Row: " + strconv.Itoa(row))
		}
		if sli[0] == "printc" {
			count, errC := strconv.Atoi(sli[1])
			if errC != nil {
				return printCommand("ERROR: The second argument has to be a number. Row: " + strconv.Itoa(row))
			}
			symbol := sli[2]
			if len(symbol) > 1 {
				return printCommand("ERROR: The third argument has to be a symbol. Row: " + strconv.Itoa(row))
			}
			return &printcCommand{count: count, symbol: symbol}
		}
		return printCommand("ERROR: This command doesn't exist. Row: " + strconv.Itoa(row))
	default:
		return printCommand("ERROR: Too many arguments. Row: " + strconv.Itoa(row))
	}
}
