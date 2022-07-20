package behavior_pattern

// Command
type Command interface {
	Execute() string
}

type DoCommand struct {
	command string
}

func NewDoCommand() *DoCommand {
	return &DoCommand{command: "execute do command"}
}

func (c *DoCommand) Execute() string {
	return c.command
}

type QueryCommand struct {
	command string
}

func NewQueryCommand() *DoCommand {
	return &DoCommand{command: "execute query command"}
}

func (c *QueryCommand) Execute() string {
	return c.command
}

type Executor struct {
	commands []Command
}

func NewExecutor(commands []Command) *Executor {
	return &Executor{commands: commands}
}

func (e *Executor) Run() string {
	result := ""
	for _, command := range e.commands {
		result += command.Execute() + "\n"
	}
	return result
}
