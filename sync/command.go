package sync

import (
	"encoding/json"
	"net/url"

	"github.com/google/uuid"
)

type CommandArgs interface {
	Type() string
}

type Command struct {
	Type   string    `json:"type"`
	Args   any       `json:"args"`
	UUID   uuid.UUID `json:"uuid"`
	TempID uuid.UUID `json:"temp_id"`
}

func NewCommand(args CommandArgs) *Command {
	return &Command{
		Type:   args.Type(),
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
	cmds := make(Commands, 0, len(args))
	for _, arg := range args {
		cmds = append(cmds, NewCommand(arg))
	}
	return cmds
}

func (cs *Commands) EncodeValues(key string, v *url.Values) error {
	b, err := json.Marshal(cs)
	if err != nil {
		return err
	}

	v.Add(key, string(b))
	return nil
}
