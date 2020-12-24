// Álvaro Castellano Vela 2020/12/24
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Card struct {
	Value int
	Next  *Card
}

type Player struct {
	ID         int
	FirstCard  *Card
	LastCard   *Card
	TotalCards int
}

func showPlayers(players []Player) {
	for _, player := range players {
		fmt.Println("Player", player.ID)
		fmt.Println("Cards (", player.TotalCards, "):")
		for card := player.FirstCard; card != nil; card = card.Next {
			fmt.Println(card.Value)
		}
		fmt.Println("")
	}
}

func status(players []Player) string {
	var result string

	for playerId, _ := range players {
		for card := players[playerId].FirstCard; card != nil; card = card.Next {
			result += strconv.Itoa(card.Value)
			result += ","
		}
		result += "|"
	}

	return result
}

func copyPlayers(origin []Player, drews []int) []Player {
	copied := make([]Player, 2)
	for playerId, _ := range origin {
		copied[playerId].ID = origin[playerId].ID
		copied[playerId].TotalCards = drews[playerId]
		orginCard := origin[playerId].FirstCard
		for counter := 0; counter < drews[playerId]; counter++ {
			copiedCard := Card{orginCard.Value, nil}
			if copied[playerId].FirstCard == nil {
				copied[playerId].FirstCard = &copiedCard
				copied[playerId].LastCard = &copiedCard
			} else {
				copied[playerId].LastCard.Next = &copiedCard
				copied[playerId].LastCard = copied[playerId].LastCard.Next
			}
			orginCard = orginCard.Next
		}
		copied[playerId].LastCard.Next = nil
	}

	return copied
}

func processFile(filename string) []Player {

	players := make([]Player, 2)
	var addedPlayers int = 0
	var player int = 0

	players[0].ID = 1
	players[1].ID = 2

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	// first Player
	scanner.Scan()
	// read cards
	for addedPlayers < 2 {
		addedPlayers++
		for scanner.Scan() {
			cardNumberString := scanner.Text()
			if cardNumberString == "" {
				player++
				scanner.Scan()
				break
			} else {
				cardNumber, _ := strconv.Atoi(cardNumberString)
				card := Card{cardNumber, nil}
				if players[player].FirstCard == nil {
					players[player].FirstCard = &card
					players[player].LastCard = &card
				} else {
					players[player].LastCard.Next = &card
					players[player].LastCard = &card
				}
				players[player].TotalCards++
			}
		}
	}

	return players
}

func play(players []Player) (int, int) {

	var result int = 0
	var winner int = -1

	var ended bool = false
	statusHistory := make(map[string]bool)

	cardsToCompare := make([]*Card, 2)
	for players[0].TotalCards != 0 && players[1].TotalCards != 0 && !ended {
		currentStatus := status(players)
		if _, ok := statusHistory[currentStatus]; ok {
			ended = true
			winner = 0
		} else {

			statusHistory[currentStatus] = true

			winner = -1
			//Draw Cards
			for playerId, _ := range players {
				cardsToCompare[playerId] = players[playerId].FirstCard
				players[playerId].FirstCard = players[playerId].FirstCard.Next
				cardsToCompare[playerId].Next = nil
				players[playerId].TotalCards--
			}
			//Compare cards
			if cardsToCompare[0].Value <= players[0].TotalCards && cardsToCompare[1].Value <= players[1].TotalCards {
				drewCards := []int{cardsToCompare[0].Value, cardsToCompare[1].Value}
				winner, _ = play(copyPlayers(players, drewCards))
			} else {
				if cardsToCompare[0].Value > cardsToCompare[1].Value {
					winner = 0
				} else {
					winner = 1
				}
			}

			//Get cards
			if players[winner].FirstCard == nil {
				players[winner].FirstCard = cardsToCompare[winner]
				players[winner].LastCard = cardsToCompare[winner]
			} else {
				players[winner].LastCard.Next = cardsToCompare[winner]
				players[winner].LastCard = players[winner].LastCard.Next
			}
			players[winner].LastCard.Next = cardsToCompare[(winner+1)%2]
			players[winner].LastCard = players[winner].LastCard.Next
			players[winner].TotalCards += 2

		}
	}

	//ended == true
	if winner == -1 { // winners was not decided by previousState
		if players[0].TotalCards == 0 {
			winner = 1
		} else {
			winner = 0
		}
	}

	for counter, card := 0, players[winner].FirstCard; card != nil; card = card.Next {
		result += card.Value * (players[winner].TotalCards - counter)
		counter++
	}

	return winner, result
}

func main() {
	var result int
	var winner int = -1
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("You must supply a file to process.")
	}
	filename := args[0]

	players := processFile(filename)
	winner, result = play(players)
	fmt.Println("Winner:", winner+1)
	fmt.Println("Result:", result)
}
