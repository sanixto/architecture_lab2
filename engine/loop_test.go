package engine

import (
	"bufio"
	"os"

	. "gopkg.in/check.v1"
)

func (s *MySuite) TestLoop(c *C) {
	cmd1 := printCommand("Hello1")
	cmd2 := printCommand("Hello2")
	cmd3 := printcCommand{count: 5, symbol: "l"}

	loop := new(Loop)
	loop.Start()
	c.Assert(false, Equals, loop.stop)
	c.Assert(0, Equals, len(loop.q.a))

	loop.Post(cmd1)
	loop.Post(cmd2)
	loop.Post(cmd3)

	c.Assert(3, Equals, len(loop.q.a))
	loop.AwaitFinish()
	c.Assert(true, Equals, loop.stop)
	c.Assert(0, Equals, len(loop.q.a))

	var txtCommands []string
	txtCommands = append(txtCommands, "print Hello1", "print Hello2", "printc 5 l")

	filename := "loop_test.txt"
	file, _ := os.Open(filename)
	defer file.Close()

	var fileRows []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cmdLine := scanner.Text()
		fileRows = append(fileRows, cmdLine)
	}

	c.Assert(len(fileRows), Equals, 3)
	for i := 0; i < 3; i++ {
		c.Assert(txtCommands[i], Equals, fileRows[i])
	}
}