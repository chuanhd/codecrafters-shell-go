package domains

import (
	"fmt"
	"os"
	"strings"
)

type EchoCommand struct{}

func (c *EchoCommand) GetName() string {
	return "echo"
}

func (c *EchoCommand) Execute(cmd Command) {
	fmt.Fprintln(os.Stdout, strings.Join(cmd.Args, " "))
}
