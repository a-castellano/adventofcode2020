// Álvaro Castellano Vela 2020/12/16
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func processFile(filename string) (map[string][]int, [][]int, []int) {

	validCodes := make(map[string][]int)
	var tickets [][]int
	var myTicket []int

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	idRule := regexp.MustCompile(`^([^:]+): ([0-9]+)-([0-9]+) or ([0-9]+)-([0-9]+)$`)

	// read rules
	for scanner.Scan() {
		match := idRule.FindAllStringSubmatch(scanner.Text(), -1)
		if len(match) != 0 {
			ranges := make([]int, 4)
			for i := 0; i < 4; i++ {
				ranges[i], _ = strconv.Atoi(match[0][i+2])
			}
			validCodes[match[0][1]] = ranges

		} else {
			break
		}
	}

	// ignore my tcket
	scanner.Scan() //your tikcets
	scanner.Scan() //your tikcets values
	myTicketSplittedString := strings.Split(scanner.Text(), ",")
	for _, stringValue := range myTicketSplittedString {
		number, _ := strconv.Atoi(stringValue)
		myTicket = append(myTicket, number)
	}

	scanner.Scan() //
	scanner.Scan() //  nearbyTickets

	for scanner.Scan() {
		splittedString := strings.Split(scanner.Text(), ",")
		var ticketValues []int
		for _, stringValue := range splittedString {
			number, _ := strconv.Atoi(stringValue)
			ticketValues = append(ticketValues, number)
		}
		if isValidTicket(validCodes, ticketValues) {
			tickets = append(tickets, ticketValues)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// add your own ticket
	tickets = append(tickets, myTicket)
	return validCodes, tickets, myTicket
}

func isValidTicket(validCodes map[string][]int, ticketValues []int) bool {
	var valid bool = true
	for _, candidate := range ticketValues {
		var validCandidate bool = false
		for _, ranges := range validCodes {
			if (candidate >= ranges[0] && candidate <= ranges[1]) || (candidate >= ranges[2] && candidate <= ranges[3]) {
				validCandidate = true
				break
			}
		}
		if validCandidate == false {
			valid = false
			break
		}
	}
	return valid
}

func getFieldPositions(validCodes map[string][]int, tickets [][]int, ticketFields int) map[string]int {

	candidatePositions := make(map[string]map[int]bool)
	positions := make(map[string]int)
	var totalCandidates int = len(validCodes)

	for position := 0; position < ticketFields; position++ {
		var field string
		for fieldName, ranges := range validCodes {
			field = fieldName
			var isthisfield bool = true
			for _, ticket := range tickets {
				if (ticket[position] < ranges[0] || ticket[position] > ranges[1]) && (ticket[position] < ranges[2] || ticket[position] > ranges[3]) {
					isthisfield = false
					break
				}
			}
			if isthisfield {
				if _, ok := candidatePositions[field]; !ok {
					newMap := make(map[int]bool)
					candidatePositions[field] = newMap
				}
				candidatePositions[field][position] = true
			}
		}
	}

	// discard duplicated inlt all fileds get only one position

	for len(positions) != totalCandidates {
		for field, candidates := range candidatePositions {
			if len(candidates) == 1 {
				// this position is unique, remove all apearen in the other candidates
				for positiontoDelete, _ := range candidates { // onlt one iteration
					for fieldToBeCleaned, _ := range candidatePositions {
						delete(candidatePositions[fieldToBeCleaned], positiontoDelete)
					}
					positions[field] = positiontoDelete
					delete(candidatePositions, field)
				}
			}
		}
	}
	return positions
}

func calculateResult(fieldPositions map[string]int, myTicket []int) int {
	var result int = 1

	for field, position := range fieldPositions {
		if strings.HasPrefix(field, "departure") {
			result = myTicket[position] * result
		}
	}

	return result
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("You must supply a file to process.")
	}
	filename := args[0]

	validCodes, tickets, myTicket := processFile(filename)
	fieldPositions := getFieldPositions(validCodes, tickets, len(myTicket))
	//	fmt.Println(validCodes, tickets, myTicket)
	result := calculateResult(fieldPositions, myTicket)
	fmt.Println("Result:", result)
}
