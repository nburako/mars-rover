package rover

import "mars-rover/src/plateau"

type Coordinate struct {
	X int
	Y int
}

type Rover struct {
	Coordinate
	Direction
	plateau.Plateau
}

func (rover *Rover) Move(movement Movement) {
	switch movement {
	case Move:
		rover.moveForward()
	case Right:
		rover.turnRight()
	case Left:
		rover.turnLeft()
	}
}

func (rover *Rover) turnLeft() {
	switch rover.Direction {
	case North:
		rover.Direction = West
	case West:
		rover.Direction = South
	case South:
		rover.Direction = East
	case East:
		rover.Direction = North
	}
}

func (rover *Rover) turnRight() {
	switch rover.Direction {
	case North:
		rover.Direction = East
	case West:
		rover.Direction = North
	case South:
		rover.Direction = West
	case East:
		rover.Direction = South
	}
}

func (rover *Rover) moveForward() {
	switch rover.Direction {
	case North:
		rover.Coordinate.Y += 1
	case West:
		rover.Coordinate.X -= 1
	case South:
		rover.Coordinate.Y -= 1
	case East:
		rover.Coordinate.X += 1
	}
}
