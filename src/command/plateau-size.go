package command

import (
	"mars-rover/src/plateau"
	"strconv"
	"strings"
)

func (c *Command) RunCommand(command string) plateau.Plateau {
	plateau := new(plateau.Plateau)
	width, height := parseCommand(command)
	plateau.Determine(width, height)

	return *plateau
}

func parseCommand(command string) (int, int) {
	result := strings.Split(command, " ")
	width, _ := strconv.ParseInt(result[0], 10, 32)
	height, _ := strconv.ParseInt(result[1], 10, 32)

	return int(width), int(height)
}
