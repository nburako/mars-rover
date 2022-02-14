package plateau

type Plateau struct {
	Size
}

func (plateau *Plateau) Determine(width int, height int) {
	plateau.width = width
	plateau.height = height
}
