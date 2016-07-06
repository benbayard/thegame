package card

type Card struct {
	Value int
}

func NewCard(value int) *Card {
	return &Card{value}
}
