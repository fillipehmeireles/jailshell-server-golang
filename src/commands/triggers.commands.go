package commands

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"tcp-server/src/environment"
	"tcp-server/src/services"
)

// Init all command's triggers
func TriggerSetup() {
	environment.SYSTEM_COMMANDS[0].Trigger = OSCommandTrigger
	environment.SYSTEM_COMMANDS[1].Trigger = HelpTrigger
	environment.SYSTEM_COMMANDS[2].Trigger = OSCommandTrigger
	environment.SYSTEM_COMMANDS[3].Trigger = OSCommandTrigger
	environment.SYSTEM_COMMANDS[4].Trigger = ChangeDirTrigger
	environment.SYSTEM_COMMANDS[5].Trigger = FileOrDirectoryPermissions
	environment.SYSTEM_COMMANDS[6].Trigger = SendFileToCient
}

// HelpTrigger is a trigger to "help" command
// Type of models.Trigger
func HelpTrigger(commandData ...string) []string {
	var systemCommands []string
	systemCommands = append(systemCommands, fmt.Sprintf("%s - %s	| %s", "ID", "Alias", "Description"))
	for _, cmd := range environment.SYSTEM_COMMANDS {
		systemCommands = append(systemCommands, fmt.Sprintf("%d - %s	| %s", cmd.ID, cmd.Alias, cmd.Description))
	}
	return systemCommands
}

// OSCommandTrigger is a trigger to "listdir" command
// Type of models.Trigger
func OSCommandTrigger(commandData ...string) []string {
	var result []string
	cmd := exec.Command(commandData[0])
	stdout, _ := cmd.Output()
	result = append(result, string(stdout))
	return result
}

// ChangeDirTrigger is a trigger to "cd" command
// Type of models.Trigger
func ChangeDirTrigger(commandData ...string) []string {
	var result []string

	dirToGo := strings.Split(commandData[0], " ")[1]
	if erro := os.Chdir(dirToGo); erro != nil {
		result = append(result, fmt.Sprintf("Error: %s", erro))
	}
	return result
}

//FileOrDirectoryPermissions is a trigger to "frperm" command
// Type of models.Trigger
func FileOrDirectoryPermissions(commandData ...string) []string {
	var fileOrDirPermissions []string

	if !strings.Contains(commandData[0], " ") {
		fileOrDirPermissions = append(fileOrDirPermissions, "No file or directory chosen. Type: help")
		return fileOrDirPermissions
	}

	file := strings.Split(commandData[0], " ")[1]

	info, erro := os.Stat(file)
	if erro != nil {
		log.Println(erro)
	}
	defer recover()

	mode := info.Mode().String()

	fileOrDirPermissions = append(fileOrDirPermissions, string(mode))

	return fileOrDirPermissions
}

func SendFileToCient(commandData ...string) []string {
	sendFileRes := services.SendFile(commandData[0], commandData[1])

	return sendFileRes
}
