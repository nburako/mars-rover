package command

import (
	"mars-rover/src/plateau"
	"mars-rover/src/rover"
	"strconv"
	"strings"
)

func (c *Command) DeployRover(command string, plateau plateau.Plateau) rover.Rover {
	coordinate, direction := parseDeployCommand(command)

	rover := rover.Rover{Coordinate: coordinate, Direction: direction, Plateau: plateau}

	return rover
}

func parseDeployCommand(command string) (rover.Coordinate, rover.Direction) {
	result := strings.Split(command, " ")
	x, _ := strconv.ParseInt(result[0], 10, 32)
	y, _ := strconv.ParseInt(result[1], 10, 32)
	d := result[2]

	var direction rover.Direction
	coordinate := rover.Coordinate{X: int(x), Y: int(y)}

	return coordinate, direction.DetermineDirection(d)
}
