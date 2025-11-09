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
	Unquoted     State = 1 << iota // 0001
	SingleQuoted                   // 0010
	DoubleQuoted                   // 0100
	Escaped                        // 1000
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

	// Special character to avoid typo
	singleQuote := '\''
	doubleQuote := '"'
	escape := '\\'
	space := ' '
	tab := '\t'

	stringAsRunes := []rune(s)
	inEscape := false
	for i := 0; i < len(stringAsRunes); i++ {
		c := stringAsRunes[i]
		switch state {
		case Unquoted:
			switch c {
			case space, tab:
				flush()
			case singleQuote:
				state = SingleQuoted // change state from Unquoted to SingleQuoted
			case doubleQuote:
				state = DoubleQuoted // change state from Unquoted to DoubleQuoted
				inEscape = false
			case escape:
				// escape next character if exists
				if i+1 < len(s) {
					buf = append(buf, stringAsRunes[i+1])
					i++ // skip next char
				} else {
					buf = append(buf, escape)
				}
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
			if inEscape {
				switch c {
				case '"', '\\':
					buf = append(buf, c)
				case '\n':
					// drop both for line-continuation
				default:
					// keep backslash literally if not one of the four
					buf = append(buf, '\\', c)
				}
				inEscape = false
				continue
			}
			switch c {
			case doubleQuote:
				state = Unquoted
			case escape:
				inEscape = true
			default:
				buf = append(buf, c)
			}
		}
	}
	if state == SingleQuoted || state == DoubleQuoted {
		return nil, fmt.Errorf("Unmatched")
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
