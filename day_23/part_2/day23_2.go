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

func processInput(input string) (Cups, map[int]*Cup) {

	cupsLocation := make(map[int]*Cup)
	cups := Cups{nil, nil}

	for _, stringNimber := range input {
		number := int(stringNimber - '0')
		cup := Cup{number, nil, nil}
		cupsLocation[number] = &cup
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
	for number := 10; number <= 1000000; number++ {
		cup := Cup{number, nil, nil}
		cupsLocation[number] = &cup
		cups.Last.Next = &cup
		cup.Previous = cups.Last
		cups.Last = cups.Last.Next
	}
	cups.Last.Next = cups.First
	cups.First.Previous = cups.Last

	return cups, cupsLocation
}

func play(cups Cups, cupsLocation map[int]*Cup, times int) int {

	var result int

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
				destination = 1000000
			}
			if _, ok := pickUpMap[destination]; !ok {
				validDestination = true
			}
			if !validDestination {
				destination--
			}
		}
		var destinationCup *Cup = cupsLocation[destination]

		pickUp.Next.Next.Next = destinationCup.Next
		destinationCup.Next.Previous = pickUp.Next.Next.Next
		pickUp.Previous = destinationCup
		destinationCup.Next = pickUp

		current = current.Next

		times--
	}

	result = cupsLocation[1].Next.Value * cupsLocation[1].Next.Next.Value
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

	cups, cupsLocation := processInput(inputString)

	result := play(cups, cupsLocation, times)

	fmt.Println("Result:", result)
}
