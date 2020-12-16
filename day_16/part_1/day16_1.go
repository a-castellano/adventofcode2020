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

func processFile(filename string) (map[string][]int, []int) {

	validCodes := make(map[string][]int)
	var ticketValues []int

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
	scanner.Scan() //
	scanner.Scan() //  nearbyTickets

	for scanner.Scan() {
		splittedString := strings.Split(scanner.Text(), ",")
		for _, stringValue := range splittedString {
			number, _ := strconv.Atoi(stringValue)
			ticketValues = append(ticketValues, number)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return validCodes, ticketValues
}

func calculateResult(validCodes map[string][]int, ticketValues []int) int {
	var result int = 0
	for _, candidate := range ticketValues {
		var valid bool = false
		for _, ranges := range validCodes {
			if (candidate >= ranges[0] && candidate <= ranges[1]) || (candidate >= ranges[2] && candidate <= ranges[3]) {
				valid = true
				break
			}
		}
		if valid == false {
			result += candidate
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

	validCodes, ticketValues := processFile(filename)
	result := calculateResult(validCodes, ticketValues)
	fmt.Println("Result", result)
}
