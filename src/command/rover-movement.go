package command

import (
	"fmt"
	"mars-rover/src/rover"
	"strconv"
)

func MoveRover(command string, r rover.Rover) {
	for i := 0; i < len(command); i++ {
		movement := new(rover.Movement)
		r.Move(movement.DetermineMovement(string(command[i])))
	}

	writeLocation(strconv.Itoa(r.Coordinate.X), strconv.Itoa(r.Coordinate.Y), string(r.Direction))
}

func writeLocation(x string, y string, direction string) {
	fmt.Printf("%s %s %s \n", x, y, direction)
}
