package engine

import "fmt"

type Command interface {
	Execute(handler Handler)
}

type printCommand string

func (pc printCommand) Execute(h Handler) {
	fmt.Println(pc)
}

type printcCommand struct {
	count  int
	symbol string
}

func (p *printcCommand) Execute(h Handler) {
	var str string
	for i := 0; i < p.count; i++ {
		str = str + p.symbol
	}
	h.Post(printCommand(str))
}

type stopCommand struct{}

func (s stopCommand) Execute(h Handler) {
	h.(*Loop).stop = true
}
