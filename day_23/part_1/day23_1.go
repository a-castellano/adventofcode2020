// Álvaro Castellano Vela 2020/12/24
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Cup struct {
	Value    int
	Next     *Cup
	Previous *Cup
}

type Cups struct {
	First *Cup
	Last  *Cup
}

func processInput(input string) (Cups, int) {

	cups := Cups{nil, nil}
	var counter int = 0

	for _, stringNimber := range input {
		counter++
		number := int(stringNimber - '0')
		cup := Cup{number, nil, nil}
		if cups.First == nil {
			cups.First = &cup
			cups.Last = &cup
			cups.First.Next = &cup
			cups.First.Previous = &cup
			cups.Last.Next = &cup
			cups.Last.Previous = &cup
		} else {
			cups.Last.Next = &cup
			cup.Previous = cups.Last
			cups.Last = cups.Last.Next
		}
	}
	cups.Last.Next = cups.First
	cups.First.Previous = cups.Last

	return cups, counter
}

func play(cups Cups, times int) string {

	var result string

	var current *Cup = cups.First

	for times > 0 {

		pickUpMap := make(map[int]bool)
		pickUpMap[current.Next.Value] = true
		pickUpMap[current.Next.Next.Value] = true
		pickUpMap[current.Next.Next.Next.Value] = true

		var pickUp *Cup = current
		pickUp = pickUp.Next
		current.Next = pickUp.Next.Next.Next
		current.Next.Previous = current

		destination := current.Value - 1
		var validDestination bool = false
		for !validDestination {
			if destination == 0 {
				destination = 9
			}
			if _, ok := pickUpMap[destination]; !ok {
				validDestination = true
			}
			if !validDestination {
				destination--
			}
		}
		var destinationCup *Cup = current
		for destinationCup.Value != destination {
			destinationCup = destinationCup.Next
		}

		pickUp.Next.Next.Next = destinationCup.Next
		destinationCup.Next.Previous = pickUp.Next.Next.Next
		pickUp.Previous = destinationCup
		destinationCup.Next = pickUp

		current = current.Next

		times--
	}
	cup := current
	for cup.Value != 1 {
		cup = cup.Next
	}
	cup = cup.Next
	for i := 0; i < 8; i++ {
		result += strconv.Itoa(cup.Value)
		cup = cup.Next
	}

	return result
}

func main() {
	//	var result int
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatal("You must supply a string to process and a number of times.")
	}
	inputString := args[0]
	times, _ := strconv.Atoi(args[1])

	cups, numberOfCups := processInput(inputString)

	cup := cups.First
	for i := 0; i < numberOfCups+1; i++ {
		fmt.Println(cup.Value)
		cup = cup.Next
	}

	result := play(cups, times)

	fmt.Println("Resul:", result)

}
