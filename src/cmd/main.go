package main

import (
	"bufio"
	c "mars-rover/src/command"
	p "mars-rover/src/plateau"
	r "mars-rover/src/rover"
	"os"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	plateau := new(p.Plateau)
	roverList := []r.Rover{}

	for {
		input, _ := inputReader.ReadString('\r')
		input = strings.TrimSuffix(input, "\r")
		input = strings.TrimPrefix(input, "\n")

		commandType := c.ResolveCommand(input)
		switch commandType {
		case c.PlateauSize:
			*plateau = commandType.RunCommand(input)
		case c.RoverDeploy:
			rover := commandType.DeployRover(input, *plateau)
			roverList = append(roverList, rover)
		case c.RoverMovement:
			c.MoveRover(input, roverList[len(roverList)-1])
		}
	}
}
