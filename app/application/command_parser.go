package application

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/domains"
)

type CommandParser struct {
	reader *bufio.Reader
}

func NewCommandParser(reader *bufio.Reader) *CommandParser {
	return &CommandParser{reader: reader}
}

func (parser *CommandParser) ParseCommand() (*domains.Command, error) {
	content, err := parser.reader.ReadString('\n')
	if err != nil {
		fmt.Fprint(os.Stderr, "ERROR: "+err.Error())
		return nil, err
	}

	// Split read content to command and its arguments
	elems := strings.Fields(content)

	return &domains.Command{
		Name: elems[0],
		Args: elems[1:],
	}, nil
}
