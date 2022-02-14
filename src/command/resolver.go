package command

import "regexp"

type Command int

const (
	PlateauSize Command = iota
	RoverDeploy
	RoverMovement
)

func ResolveCommand(command string) Command {
	isPlateauSize, _ := regexp.MatchString("^\\d+ \\d+$", command)

	if isPlateauSize {
		return PlateauSize
	}

	isRoverDeploy, _ := regexp.MatchString("^\\d+ \\d+ [ESWN]$", command)

	if isRoverDeploy {
		return RoverDeploy
	}

	isRoverMovement, _ := regexp.MatchString("^[LRM]+$", command)

	if isRoverMovement {
		return RoverMovement
	}

	panic("Invalid command!")
}
