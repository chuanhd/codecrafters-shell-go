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

const (
	CharSingleQuote = '\''
	CharDoubleQuote = '"'
	CharEscape      = '\\'
	CharSpace       = ' '
	CharTab         = '\t'
)

func extractCommandAndArgs(content string) (cmd, argumentStr string) {
	if len(content) == 0 {
		return "", ""
	}

	first := rune(content[0])
	if first != CharSingleQuote && first != CharDoubleQuote {
		cmd, argumentStr, _ = strings.Cut(content, " ")
		return
	}

	endIdx := strings.IndexRune(content[1:], first)
	if endIdx == -1 {
		cmd, argumentStr, _ = strings.Cut(content, " ")
		return
	}

	endIdx += 1
	cmd = content[1:endIdx]
	argumentStr = content[endIdx+1:]
	return
}

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

	stringAsRunes := []rune(s)
	inEscape := false
	for i := 0; i < len(stringAsRunes); i++ {
		c := stringAsRunes[i]
		switch state {
		case Unquoted:
			switch c {
			case CharSpace, CharTab:
				flush()
			case CharSingleQuote:
				state = SingleQuoted // change state from Unquoted to SingleQuoted
			case CharDoubleQuote:
				state = DoubleQuoted // change state from Unquoted to DoubleQuoted
				inEscape = false
			case CharEscape:
				// escape next character if exists
				if i+1 < len(s) {
					buf = append(buf, stringAsRunes[i+1])
					i++ // skip next char
				} else {
					buf = append(buf, CharEscape)
				}
			default:
				buf = append(buf, c) // Append normal character to buffer
			}
		case SingleQuoted:
			if c == CharSingleQuote {
				state = Unquoted // Closed single quoted. Change state from SingleQuoted to Unquoted
			} else {
				buf = append(buf, c) // Inside single quote, append character to buffer
			}
		case DoubleQuoted:
			if inEscape {
				switch c {
				case CharDoubleQuote, CharEscape:
					buf = append(buf, c)
				default:
					// keep backslash literally if not one of the four
					buf = append(buf, '\\', c)
				}
				inEscape = false
				continue
			}
			switch c {
			case CharDoubleQuote:
				state = Unquoted
			case CharEscape:
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

func parseCmdAndRedirectArgs(rawArgs []string) (*domains.RedirectArgument, []string, error) {
	outPath, errPath := "", ""
	needAppend := false
	var args []string
	for i := 0; i < len(rawArgs); i++ {
		switch rawArgs[i] {
		case ">", "1>":
			if i+1 > len(rawArgs) {
				return nil, args, fmt.Errorf("syntax error: missing filename after %q", rawArgs[i])
			}
			outPath = rawArgs[i+1]
			needAppend = false
			i++ // skip filename
		case ">>", "1>>":
			if i+1 > len(rawArgs) {
				return nil, args, fmt.Errorf("syntax error: missing filename after %q", rawArgs[i])
			}
			outPath = rawArgs[i+1]
			needAppend = true
			i++ // skip filename
		case "2>":
			if i+1 > len(rawArgs) {
				return nil, args, fmt.Errorf("syntax error: missing filename after %q", rawArgs[i])
			}
			errPath = rawArgs[i+1]
			needAppend = false
			i++ // skip filename
		case "2>>":
			if i+1 > len(rawArgs) {
				return nil, args, fmt.Errorf("syntax error: missing filename after %q", rawArgs[i])
			}
			errPath = rawArgs[i+1]
			needAppend = true
			i++ // skip filename
		default:
			args = append(args, rawArgs[i])
		}
	}

	return &domains.RedirectArgument{
		StdOutPath:   outPath,
		StdErrPath:   errPath,
		StdOutAppend: needAppend,
	}, args, nil
}

func (parser *CommandParser) ParseCommand() (*domains.Command, error) {
	content, err := parser.reader.ReadString('\n')
	if err != nil {
		fmt.Fprint(os.Stderr, "ERROR: "+err.Error())
		return nil, err
	}

	cmd, argumentStr := extractCommandAndArgs(strings.TrimSpace(content))
	args, err := tokenize(argumentStr)
	redirectArgs, cmdArgs, err := parseCmdAndRedirectArgs(args)

	if err != nil {
		fmt.Fprint(os.Stderr, "ERROR: "+err.Error())
		return nil, err
	}

	return &domains.Command{
		Name:        cmd,
		Args:        cmdArgs,
		Writer:      os.Stdout,
		ErrWriter:   os.Stderr,
		RedirectArg: *redirectArgs,
	}, nil
}
