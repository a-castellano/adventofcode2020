// Álvaro Castellano Vela 2020/12/15
package main

import (
	//	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func processInput(inputString string) []int {
	var input []int

	splittedString := strings.Split(inputString, ",")

	for _, stringNumber := range splittedString {
		number, _ := strconv.Atoi(stringNumber)
		input = append(input, number)
	}

	return input
}

func play(input []int, times int) int {

	var lastNumber int
	var turn int = 1
	numberSpoken := make(map[int]bool)
	numberTurn := make(map[int][]int)

	for _, number := range input {
		lastNumber = number
		numberTurn[lastNumber] = []int{turn, 0}
		turn++
	}
	for turn <= times {
		if _, ok := numberSpoken[lastNumber]; !ok {
			lastNumber = 0
			numberSpoken[lastNumber] = true
		} else {
			difference := numberTurn[lastNumber][0] - numberTurn[lastNumber][1]
			lastNumber = difference
			numberSpoken[lastNumber] = true
		}
		if _, ok := numberTurn[lastNumber]; !ok {
			numberTurn[lastNumber] = []int{turn, 0}
		}
		numberTurn[lastNumber][1] = numberTurn[lastNumber][0]
		numberTurn[lastNumber][0] = turn
		turn++
	}

	return lastNumber
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("You must supply a string to process.")
	}
	inputString := args[0]

	input := processInput(inputString)

	finalNumber := play(input, 30000000)

	fmt.Println("Result:", finalNumber)
}
