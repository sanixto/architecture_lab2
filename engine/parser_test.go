package engine

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestParse(c *C) {
	cmd := Parse("print Hello")
	c.Assert(printCommand("Hello"), Equals, cmd)

	cmd = Parse("printc 2 n")
	c.Assert(printcCommand{count: 2, symbol: "n"}, Equals, cmd)

	cmd = Parse("")
	c.Assert(printCommand("ERROR: The empty string."), Equals, cmd)

	cmd = Parse(" ")
	c.Assert(printCommand("ERROR: The empty string."), Equals, cmd)

	cmd = Parse("print")
	c.Assert(printCommand("ERROR: No arguments."), Equals, cmd)

	cmd = Parse("printc")
	c.Assert(printCommand("ERROR: No arguments."), Equals, cmd)

	cmd = Parse("printc 5")
	c.Assert(printCommand("ERROR: Too few arguments."), Equals, cmd)

	cmd = Parse("print 2efwfw 3")
	c.Assert(printCommand("ERROR: Too many arguments."), Equals, cmd)

	cmd = Parse("printc 2 l 4")
	c.Assert(printCommand("ERROR: Too many arguments."), Equals, cmd)

	cmd = Parse("printc g l")
	c.Assert(printCommand("ERROR: The second argument has to be a number."), Equals, cmd)

	cmd = Parse("printc 3 4fwafe")
	c.Assert(printCommand("ERROR: The third argument has to be a symbol."), Equals, cmd)

	cmd = Parse("someCommand 3 0")
	c.Assert(printCommand("ERROR: This command doesn't exist."), Equals, cmd)
}