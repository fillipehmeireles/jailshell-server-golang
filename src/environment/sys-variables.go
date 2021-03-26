package environment

import (
	"tcp-server/src/models"
)

// All system commands
var SYSTEM_COMMANDS = make([]models.Command, 7)

// All command variable Prefixes
var VAR_PREFIXES = make([]string, 2)

// Init All variables
func VariablesSetup() {
	VAR_PREFIXES = []string{"cd", "frperm"}

	SYSTEM_COMMANDS[0] = models.Command{
		ID:          0,
		Alias:       "listdir",
		IsOSCMD:     true,
		UnixCommand: "ls",
		HasArgs:     false,
		Description: "All files and directories on current dir",
	}
	SYSTEM_COMMANDS[1] = models.Command{
		ID:          1,
		Alias:       "help",
		IsOSCMD:     false,
		UnixCommand: "",
		HasArgs:     false,
		Description: "All commands",
	}

	SYSTEM_COMMANDS[2] = models.Command{
		ID:          2,
		Alias:       "currdir",
		IsOSCMD:     true,
		UnixCommand: "pwd",
		HasArgs:     false,
		Description: "Working Directory",
	}

	SYSTEM_COMMANDS[3] = models.Command{
		ID:          3,
		Alias:       "clear",
		IsOSCMD:     true,
		UnixCommand: "clear",
		HasArgs:     false,
		Description: "Clear Screen",
	}

	SYSTEM_COMMANDS[4] = models.Command{
		ID:          4,
		Alias:       "cd",
		IsOSCMD:     true,
		UnixCommand: "cd",
		HasArgs:     true,
		Description: "Change Directory",
	}

	SYSTEM_COMMANDS[5] = models.Command{
		ID:          5,
		Alias:       "frperm",
		IsOSCMD:     false,
		UnixCommand: "",
		HasArgs:     true,
		Description: "File or Directory Permissions. Use: frperm <file or directory name>",
	}

	SYSTEM_COMMANDS[6] = models.Command{
		ID:          6,
		Alias:       "getserverf",
		IsOSCMD:     false,
		UnixCommand: "",
		HasArgs:     true,
		Description: "Download file from server. Use: getserverf <file name>",
	}
}
