package path

import (
	"errors"
	"fmt"
	"strings"
)

//CommandPath command path
type CommandPath struct {
	CommandName string
	Domain      string
	Subdomain   string
}

//ErrUnknownCommand unknown command error
var ErrUnknownCommand = errors.New("unknown command")

//ParseCommand parses command
func ParseCommand(commandText string) (CommandPath, error) {
	commandParts := strings.SplitN(commandText, "__", 3)
	if len(commandParts) != 3 {
		return CommandPath{}, ErrUnknownCommand
	}

	return CommandPath{
		CommandName: commandParts[0],
		Domain:      commandParts[1],
		Subdomain:   commandParts[2],
	}, nil
}

//WithCommandName sets command name
func (c CommandPath) WithCommandName(commandName string) CommandPath {
	c.CommandName = commandName

	return c
}

func (c CommandPath) String() string {
	return fmt.Sprintf("/%s__%s__%s", c.CommandName, c.Domain, c.Subdomain)
}
