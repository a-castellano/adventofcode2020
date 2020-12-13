// Álvaro Castellano Vela 2020/12/13
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
		}
	}
	busIDs = append(busIDs, currentId)

	sort.Ints(busIDs)

	return timestamp, busIDs
}

func calculateDepart(timestamp int, busIDs []int) int {
	var result int = -1
	var closerValue int = 999999999
	var closerID int = -1

	for _, ID := range busIDs {
		var time int = 0
		for time < timestamp {
			time += ID
		}
		if time < closerValue {
			closerValue = time
			closerID = ID
		}
	}

	result = (closerValue - timestamp) * closerID

	return result
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("You must supply a file to process.")
	}
	filename := args[0]

	timestamp, busIDs := processFile(filename)
	result := calculateDepart(timestamp, busIDs)

	fmt.Println("Result:", result)
}
