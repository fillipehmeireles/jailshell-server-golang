package syshandler

import (
	"fmt"
	"log"
	"net"
	"tcp-server/src/models"
)

type Commands struct {
	*models.Command
}

// Struct instance to add new methods
func OperateCommandCore(cmd *models.Command) *Commands {
	return &Commands{cmd}
}

// Run Command Operation
func (coreCommand *Commands) RunCommand(c net.Conn, args ...string) {
	var results []string
	var arg string
	if coreCommand.HasArgs {
		arg = args[0]
	} else {
		arg = coreCommand.UnixCommand
	}

	result, err := coreCommand.Trigger(arg)
	if err != nil {
		results = append(results, err.Error())
	} else {
		results = result
	}
	log.Printf("%s -> %s", c.RemoteAddr().String(), coreCommand.Alias)
	for _, result := range results {
		c.Write([]byte(fmt.Sprintln(result)))
	}
}
