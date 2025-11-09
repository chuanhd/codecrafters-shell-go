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

type State int

const (
	Unquoted State = iota
	SingleQuoted
	DoubleQuoted
)

func tokenize(s string) ([]string, error) {
	state := Unquoted
	var tok []string
	var buf []rune

	flush := func() {
		if len(buf) > 0 {
			tok = append(tok, string(buf))
			buf = nil
		}
	}

	singleQuote := '\''
	doubleQuote := '"'
	for _, c := range s {
		switch state {
		case Unquoted:
			switch c {
			case ' ', '\t':
				flush()
			case singleQuote:
				state = SingleQuoted // change state from Unquoted to SingleQuoted
			case doubleQuote:
				state = DoubleQuoted // change state from Unquoted to DoubleQuoted
			default:
				buf = append(buf, c) // Append normal character to buffer
			}
		case SingleQuoted:
			if c == singleQuote {
				state = Unquoted // Closed single quoted. Change state from SingleQuoted to Unquoted
			} else {
				buf = append(buf, c) // Inside single quote, append character to buffer
			}
		case DoubleQuoted:
			if c == doubleQuote {
				state = Unquoted // Closed double quoted. Change state from DoubleQuoted to Unquoted
			} else {
				buf = append(buf, c) // Inside single quote, append character to buffer
			}
		}
	}
	if state == SingleQuoted {
		return nil, fmt.Errorf("Unmached '")
	}
	flush()

	return tok, nil
}

func (parser *CommandParser) ParseCommand() (*domains.Command, error) {
	content, err := parser.reader.ReadString('\n')
	if err != nil {
		fmt.Fprint(os.Stderr, "ERROR: "+err.Error())
		return nil, err
	}

	cmd, argumentStr, _ := strings.Cut(strings.TrimSpace(content), " ")
	args, err := tokenize(argumentStr)

	if err != nil {
		fmt.Fprint(os.Stderr, "ERROR: "+err.Error())
		return nil, err
	}

	return &domains.Command{
		Name: cmd,
		Args: args,
	}, nil
}
