package models

// Trigger is a command trigger function
type Trigger func(commandData ...string) []string

// Command is a model of commands
type Command struct {
	ID          uint32
	Alias       string
	IsOSCMD     bool
	UnixCommand string
	Description string
	HasArgs     bool
	Trigger     Trigger
}
