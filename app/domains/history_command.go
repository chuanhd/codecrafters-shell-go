package domains

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/infra"
	"github.com/codecrafters-io/shell-starter-go/app/utils"
)

type HistoryCommand struct {
	history infra.HistoryStore
}

func NewHistoryCommand(history infra.HistoryStore) *HistoryCommand {
	return &HistoryCommand{
		history: history,
	}
}

func (cmd *HistoryCommand) GetName() string {
	return "history"
}

func (c *HistoryCommand) Execute(cmd *Command) {
	if len(cmd.Args) > 1 {
		switch cmd.Args[0] {
		case "-r":
			if len(cmd.Args) < 2 {
				fmt.Fprintln(cmd.Writer, "history: -r requires a file path")
				return
			}

			path := cmd.Args[1]
			lines, err := c.readHistoryFile(path)
			if err != nil {
				fmt.Fprintf(cmd.Writer, "history: failed to read file '%s': %v\n", path, err)
				return
			}

			for _, line := range lines {
				c.history.Add(line)
			}
		case "-w":
			if len(cmd.Args) < 2 {
				fmt.Fprintln(cmd.Writer, "history: -w requires a file path")
				return
			}

			path := cmd.Args[1]
			c.writeHistoryFile(path, false)
		}

		return
	}

	var total = len(c.history.List())
	var limit = total
	if len(cmd.Args) > 0 {
		if limitArg, err := strconv.Atoi(cmd.Args[0]); err == nil {
			limit = limitArg
		}
	}

	for i := total - limit; i < total; i++ {
		line := c.history.Get(i)
		fmt.Fprintf(cmd.Writer, "%d  %s\n", i+1, line)
	}
}

func (c *HistoryCommand) readHistoryFile(path string) ([]string, error) {
	content, err := utils.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(content, "\n")

	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	return lines, nil
}

func (c *HistoryCommand) writeHistoryFile(path string, needAppend bool) {
	file, err := utils.OpenFile(path, needAppend)

	if err != nil {
		return
	}

	defer file.Close()

	content := strings.Join(c.history.List(), "\n")
	if content != "" {
		content += "\n"
	}

	_, err = file.WriteString(content)
}
