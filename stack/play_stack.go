package stack

import "github.com/benbayard/thegame/card"

type Direction int

const (
	INCREMENT Direction = 1
	DECREMENT Direction = -1
)

type PlayStack struct {
	cards     []*card.Card
	Direction Direction
}

func NewPlayStack(direction Direction) *PlayStack {
	stack := PlayStack{
		cards:     make([]*card.Card, 1),
		Direction: direction,
	}
	switch direction {
	case INCREMENT:
		stack.cards[0] = card.NewCard(1)
		break
	case DECREMENT:
		stack.cards[0] = card.NewCard(100)
		break
	default:
		panic(direction)
	}
	return &stack
}

func (p *PlayStack) CurrentCard() *card.Card {
	return p.cards[len(p.cards)-1]
}

func (p *PlayStack) DifferenceFromCurrentCard(card card.Card) int {
	return p.CurrentCard().Value - card.Value
}

func (p *PlayStack) AppendCard(card card.Card) {
	if !p.CanPlayCard(card) {
		panic("Cannot play this card, it should not have gotten here")
	}
	p.cards = append(p.cards, &card)
}

func (p *PlayStack) CanPlayCard(card card.Card) bool {
	valueDifference := p.DifferenceFromCurrentCard(card)
	if valueDifference == 0 {
		panic("Value Difference was 0, this should not happen")
	}
	switch p.Direction {
	case INCREMENT:
		return valueDifference < 0 || valueDifference == 10
	case DECREMENT:
		return valueDifference > 0 || valueDifference == -10
	}
	return false
}
