// Álvaro Castellano Vela 2020/12/13
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func processFile(filename string) (int, []int) {

	var timestamp int
	var busIDs []int

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	timestamp, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	busIDsString := []rune(scanner.Text())
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var currentId int = 0
	var currentIdPosition = 1
	var readedNumber bool = false

	for i := len(busIDsString) - 1; i >= 0; i-- {
		character := busIDsString[i]
		if character >= '0' && character <= '9' {
			currentId += (int(character) - '0') * currentIdPosition
			currentIdPosition *= 10
			readedNumber = true
		} else if character == ',' && readedNumber == true {
			busIDs = append(busIDs, currentId)
			currentId = 0
			currentIdPosition = 1
			readedNumber = false
		} else if character == 'x' {
			busIDs = append(busIDs, -1)
		}
	}
	busIDs = append(busIDs, currentId)
	// reverse array
	for i, j := 0, len(busIDs)-1; i < j; i, j = i+1, j-1 {
		busIDs[i], busIDs[j] = busIDs[j], busIDs[i]
	}

	return timestamp, busIDs
}

func calculateDepart(busIDs []int) int {
	var found bool = false
	timestamps := make([]int, len(busIDs))

	var maxID int = -1
	var maxIDPos int = -1
	//check maximun ID
	for i, ID := range busIDs {
		if ID > maxID {
			maxID = ID
			maxIDPos = i
		}
	}
	for found == false {
		timestamps[maxIDPos] += busIDs[maxIDPos]

		found = true
		for i := 0; i < len(busIDs); i++ {
			if busIDs[i] != -1 {
				difference := i - maxIDPos
				if (timestamps[maxIDPos]+difference)%busIDs[i] != 0 {
					found = false
					break
				}
			}
		}
	}

	return timestamps[maxIDPos] - maxIDPos
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("You must supply a file to process.")
	}
	filename := args[0]

	_, busIDs := processFile(filename)

	depart := calculateDepart(busIDs)
	fmt.Println("Depart:", depart)
}
