package engine

import (
	"fmt"
	"strconv"
	"strings"
)

func Parse(cmdLine string) (Command, error) {
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
