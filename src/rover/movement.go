package rover

type Movement string

const (
	Move  Movement = "M"
	Right Movement = "R"
	Left  Movement = "L"
)

func (m *Movement) DetermineMovement(s string) Movement {
	switch s {
	case "M":
		return Move
	case "R":
		return Right
	case "L":
		return Left
	}
	panic("Unknown movement!")
}
