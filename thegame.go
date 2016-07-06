package main

import "fmt"
import "github.com/benbayard/thegame/stack"
import (
	"github.com/benbayard/thegame/deck"
	"github.com/benbayard/thegame/game"
	"github.com/benbayard/thegame/player"
)

var (
	MAX_HAND_SIZE_PER_PLAYER_COUNT = map[int]int{
		2: 8,
		3: 7,
		4: 7,
		5: 6,
	}
)

const (
	INITIAL_DECK_SIZE = 98
	NUM_STACKS        = 4
	NUM_PLAYERS       = 4
)

func main() {
	gameNotYetWon := true
	numGames := 0
	gameWins := make([]int, 0)
	for gameNotYetWon {
		numGames += 1
		players := make([]*player.Player, NUM_PLAYERS)
		for index := range players {
			players[index] = &player.Player{Order: index + 1}
		}

		deck := deck.NewDeck(INITIAL_DECK_SIZE, players, MAX_HAND_SIZE_PER_PLAYER_COUNT[NUM_PLAYERS])
		stacks := make([]*stack.PlayStack, NUM_STACKS)

		for i := range stacks {
			direction := stack.INCREMENT

			if i%2 == 0 {
				direction = stack.DECREMENT
			}

			stacks[i] = stack.NewPlayStack(direction)
		}

		game := game.Game{
			Players: players,
			Deck:    &deck,
			Stacks:  stacks,
		}

		numTurns := 0

		for !game.Won {
			err := game.TakeTurn()

			numTurns += 1

			if err != nil {
				//fmt.Printf("\nNum Turns before loss: %v\n", numTurns)
				break
			}
		}

		if game.Won {
			//fmt.Printf("\nYou Won! It took %v games", numGames)
			gameWins = append(gameWins, numGames)
			if len(gameWins) == 10000 {
				gameNotYetWon = false
			}
			numGames = 0
		}
	}

	sum := 0
	maxNum := 0

	for _, val := range gameWins {
		if val > maxNum {
			maxNum = val
		}
		sum += val
	}

	fmt.Printf("\n\nThe average number of games it takes to win is: %v\nIt took %v games at most", sum/len(gameWins), maxNum)
}
