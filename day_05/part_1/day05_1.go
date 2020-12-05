// Álvaro Castellano Vela 2020/12/05
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func processFile(filename string) int {

	var highestSeatID int = -1

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		var minColumn int = 0
		var maxColumn int = 127
		var minSeat int = 0
		var maxSeat int = 7

		for _, character := range line[:7] {
			if character == 'F' {
				maxColumn -= (maxColumn-minColumn)/2 + 1
			} else {

				minColumn += (maxColumn-minColumn)/2 + 1
			}
		}
		for _, character := range line[7:] {
			if character == 'L' {
				maxSeat -= (maxSeat-minSeat)/2 + 1
			} else {

				minSeat += (maxSeat-minSeat)/2 + 1
			}
		}
		var id = maxColumn*8 + maxSeat
		if id > highestSeatID {
			highestSeatID = id
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return highestSeatID
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("You must supply a file to process.")
	}
	filename := args[0]
	highestSeatID := processFile(filename)

	fmt.Println("Highest seat ID: ", highestSeatID)
}
