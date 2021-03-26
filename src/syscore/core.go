package syscore

import (
	"errors"

	"tcp-server/src/commands"
	"tcp-server/src/environment"
	"tcp-server/src/models"
)

// Setup Commands and Triggers of System
func SetupCore() {
	environment.VariablesSetup()
	commands.TriggerSetup()
}

// Validate existence of command on system core
func ValidateCommand(command string) (models.Command, error) {
	for _, cmd := range environment.SYSTEM_COMMANDS {
		if command == cmd.Alias {
			return cmd, nil
		}
	}
	return models.Command{}, errors.New("command not found")
}
