package deck

import (
	"github.com/benbayard/thegame/card"
	"github.com/benbayard/thegame/player"
	"math/rand"
	"time"
)

type Deck struct {
	Cards []*card.Card
}

func (d Deck) Shuffle() Deck {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 1; i < len(d.Cards); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			d.Cards[r], d.Cards[i] = d.Cards[i], d.Cards[r]
		}
	}

	return d
}

func NewDeck(deckSize int, players []*player.Player, handSize int) Deck {
	cards := make([]*card.Card, deckSize)
	for index := range cards {
		cards[index] = card.NewCard(index + 2)
	}

	deck := Deck{cards}.Shuffle()
	for _, player := range players {
		deck.DrawCards(player, handSize)
	}

	return deck
}

func (deck *Deck) DrawCards(p *player.Player, num int) {
	hand, newCards := deck.Cards[0:num], deck.Cards[num:]
	deck.Cards = newCards

	oldCards := make([]*card.Card, 0)

	if p.Hand != nil {
		oldCards = p.Hand.Cards
	}

	p.Hand = &player.Hand{}

	p.Hand.Cards = append(p.Hand.Cards, append(oldCards, hand...)...)
}
