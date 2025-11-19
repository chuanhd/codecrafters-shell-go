package domains

import (
	"fmt"
	"strings"
)

type EchoCommand struct{}

func (c *EchoCommand) GetName() string {
	return "echo"
}

func (c *EchoCommand) Execute(cmd *Command) {
	fmt.Fprintln(cmd.Writer, strings.Join(cmd.Args, " "))
}
