package domains

type Command struct {
	Name string
	Args []string
}

type CommandExecutor interface {
	GetName() string
	Execute(cmd Command)
}
