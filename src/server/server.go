package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"tcp-server/src/environment"
	"tcp-server/src/models"
	"tcp-server/src/syscore"
	"tcp-server/src/syshandler"
)

// HandleConnection manages TCP connections
func HandleConnection(c net.Conn) {
	clientRemoteAddr := c.RemoteAddr().String()
	log.Println(fmt.Sprintf("Client %s connected ", clientRemoteAddr))
	for {
		c.Write([]byte(fmt.Sprintf("%s > ", clientRemoteAddr)))
		clientData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Println(fmt.Sprintf("Client %s left server", clientRemoteAddr))
				break
			}
		}
		resp := strings.TrimSuffix(clientData, "\n")
		if resp == "quit" {
			log.Println(fmt.Sprintf("Client %s left server", clientRemoteAddr))
			break
		}
		var commandOnInput string = resp

		for _, prefix := range environment.VAR_PREFIXES {
			if strings.HasPrefix(resp, prefix) {
				commandOnInput = strings.Split(resp, " ")[0]
			}
		}

		var commandCoreCheck models.Command

		if resp != "" {
			commandCoreCheck, err = syscore.ValidateCommand(commandOnInput)
			if err != nil {
				c.Write([]byte(fmt.Sprintf("%s\n", err)))
			} else {
				sysExec := syshandler.OperateCommandCore(&commandCoreCheck)
				sysExec.RunCommand(c, resp)
			}
		}
	}
}
