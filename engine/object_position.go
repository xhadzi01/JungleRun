package engine

const scale float64 = 100.0

type Position int

func (p Position) Value() float64 {

	return float64(p) / scale
}

func (pos Position) ValueFloored() int {
	posCasted := int(pos)
	scaleCasted := int(scale)

	floored := posCasted / scaleCasted
	if floored*scaleCasted == posCasted || posCasted >= 0 {
		return floored
	}
	return floored - 1
}

type ObjectPosition struct {
	xPos    Position
	yPos    Position
	gravity Position
}
