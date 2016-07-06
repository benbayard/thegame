package game

import (
	//"fmt"
	"github.com/benbayard/thegame/deck"
	"github.com/benbayard/thegame/player"
	"github.com/benbayard/thegame/stack"
)

const (
	CARDS_TO_PLAY = 2
)

type Game struct {
	Deck      *deck.Deck
	Players   []*player.Player
	Stacks    []*stack.PlayStack
	Won       bool
	CantBeWon bool
}

func (g *Game) TakeTurn() error {
	//fmt.Printf("\nSTART OF THE TURN\n")
	for _, player := range g.Players {
		//fmt.Printf("\nPlayer: %v\n", player.Order)
		if len(player.Hand.Cards) == 0 {
			break
		}
		cardsToPlay := CARDS_TO_PLAY
		if len(g.Deck.Cards) == 0 {
			cardsToPlay = 1
		}
		cardsPlayed, err := player.PlayCards(g.Stacks, cardsToPlay)
		if err != nil {
			g.CantBeWon = true
			return err
		}

		//fmt.Printf("\nEnd Of Player %v Round: \n", player.Order)
		//for _, stacky := range g.Stacks {
		//	if stacky.Direction == stack.INCREMENT {
		//		fmt.Printf("%v↑ ", stacky.CurrentCard().Value)
		//	} else {
		//		fmt.Printf("%v↓ ", stacky.CurrentCard().Value)
		//	}
		//}

		if len(g.Deck.Cards) != 0 {
			g.Deck.DrawCards(player, cardsPlayed)
		}
		//fmt.Printf("\nHand Size: %v, Deck Size: %v\n", len(player.Hand.Cards), len(g.Deck.Cards))
	}

	if len(g.Deck.Cards) == 0 {
		gameIsWon := true
		for _, player := range g.Players {
			if len(player.Hand.Cards) != 0 && gameIsWon {
				gameIsWon = false
			}
		}

		g.Won = gameIsWon
	}
	return nil
}
