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
	if coreCommand.HasArgs {
		results = coreCommand.Trigger(args[0])
	} else {
		results = coreCommand.Trigger(coreCommand.UnixCommand)
	}

	log.Println(fmt.Sprintf("%s -> %s", c.RemoteAddr().String(), coreCommand.Alias))
	for _, result := range results {
		c.Write([]byte(fmt.Sprintln(result)))
	}
}
