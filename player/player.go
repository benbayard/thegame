package player

import (
	//"fmt"
	"errors"
	"github.com/benbayard/thegame/card"
	"github.com/benbayard/thegame/stack"
	"math"
)

type Player struct {
	Hand  *Hand
	Order int
}

type Hand struct {
	Cards    []*card.Card
	handSize int
}

type CardStackPair struct {
	card  *card.Card
	stack *stack.PlayStack
	index int
}

func (p *Player) PlayCards(stacks []*stack.PlayStack, numCards int) (int, error) {
	for i := 0; i < numCards; i++ {
		card, stack, index := p.BestCardToPlay(stacks)
		if card == nil {
			return i, errors.New("No valid play")
		}
		//for _, card := range p.Hand.Cards {
		//	fmt.Printf("%v ", card.Value)
		//}
		//fmt.Printf("Card Played: %v\n", card.Value)
		p.Hand.Cards = append(p.Hand.Cards[:index], p.Hand.Cards[index+1:]...)
		stack.AppendCard(*card)
	}
	return numCards, nil
}

func (p *Player) BestCardToPlay(stacks []*stack.PlayStack) (*card.Card, *stack.PlayStack, int) {
	differences := make(map[int]CardStackPair)
	for _, stacky := range stacks {
		for index, card := range p.Hand.Cards {
			if stacky.CanPlayCard(*card) {
				difference := stacky.DifferenceFromCurrentCard(*card)
				if (stacky.Direction == stack.INCREMENT && difference == 10) ||
					(stacky.Direction == stack.DECREMENT && difference == -10) {
					difference = -10
				} else {
					difference = int(math.Abs(float64(difference)))
				}

				differences[difference] = CardStackPair{card, stacky, index}
			}
		}
	}

	if len(differences) == 0 {
		//fmt.Printf("\nHand:  ")
		//for _, card := range p.Hand.Cards {
		//	fmt.Printf("%v ", card.Value)
		//}
		//fmt.Printf("\nStack: ")
		//for _, stacky := range stacks {
		//	if stacky.Direction == stack.INCREMENT {
		//		fmt.Printf("%v↑ ", stacky.CurrentCard().Value)
		//	} else {
		//		fmt.Printf("%v↓ ", stacky.CurrentCard().Value)
		//	}
		//}
		//fmt.Printf("\n\n")
		return nil, nil, 0
	}

	currentDifference := 100
	bestCardStack := CardStackPair{}

	for difference, cardStack := range differences {
		if currentDifference != -10 && (difference < currentDifference || difference == -10) {
			currentDifference = difference
			bestCardStack = cardStack
		}
	}

	return bestCardStack.card, bestCardStack.stack, bestCardStack.index
}
