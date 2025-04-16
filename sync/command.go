package sync

import "github.com/google/uuid"

type CommandArgs interface {
	command() string
}

type Command struct {
	// The type of the command.
	Type string `json:"type"`

	// The parameters of the command.
	Args any `json:"args"`

	// Command UUID. More details about this below.
	UUID uuid.UUID `json:"uuid"`

	// Temporary resource ID, Optional. Only specified for commands that create a
	// new resource (e.g. item_add command).
	TempID uuid.UUID `json:"temp_id"`
}

func NewCommand(args CommandArgs) *Command {
	return &Command{
		Type:   args.command(),
		Args:   args,
		UUID:   uuid.New(),
		TempID: uuid.New(),
	}
}

func NewCommandWithTempID(args CommandArgs, tempID uuid.UUID) *Command {
	c := NewCommand(args)
	c.TempID = tempID
	return c
}

type Commands []*Command

func NewCommands[T CommandArgs](args []T) Commands {
	cmds := make(Commands, len(args))
	for i, arg := range args {
		cmds[i] = NewCommand(arg)
	}
	return cmds
}
