package commands_test

import (
	"tcp-server/src/commands"
	"testing"
)

func TestCommands(t *testing.T) {
	//commands.OSCommandTrigger("pwd")

	commands.FileOrDirectoryPermissions("frperm")
}
