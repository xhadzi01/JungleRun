package engine

const scale float64 = 100.0

type Position float64

func (p *Position) Value() float64 {
	return float64(*p) / scale
}

func (pos *Position) ValueFloored() int {
	posCasted := int(*pos)
	scaleCasted := int(scale)

	floored := posCasted / scaleCasted
	if floored*scaleCasted == posCasted || posCasted >= 0 {
		return floored
	}
	return floored - 1
}

func (p *Position) Reset(value float64) {
	*p = Position(value * scale)
}

func (p *Position) Add(addition float64) {
	*p = Position(float64(*p) + addition*scale)
}

type ObjectPosition struct {
	xPos    Position
	yPos    Position
	gravity Position
}
