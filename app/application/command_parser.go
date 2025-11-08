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
	for _, c := range s {
		switch state {
		case Unquoted:
			switch c {
			case ' ', '\t':
				flush()
			case singleQuote:
				state = SingleQuoted
			default:
				buf = append(buf, c)
			}
		case SingleQuoted:
			if c == singleQuote {
				state = Unquoted
			} else {
				buf = append(buf, c)
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
