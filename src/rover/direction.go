package rover

type Direction string

const (
	North Direction = "N"
	South Direction = "S"
	East  Direction = "E"
	West  Direction = "W"
)

func (d *Direction) DetermineDirection(s string) Direction {
	switch s {
	case "N":
		return North
	case "S":
		return South
	case "E":
		return East
	case "W":
		return West
	}

	panic("Unknown direction!")
}
