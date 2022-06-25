package engine

import "fmt"

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
