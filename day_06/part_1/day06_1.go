// Álvaro Castellano Vela 2020/12/06
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func processFile(filename string) []map[rune]bool {

	groupsAnswers := make([]map[rune]bool, 0)

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	groupAnswers := make(map[rune]bool)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			for _, answer := range line {
				groupAnswers[answer] = true
			}
		} else {
			groupsAnswers = append(groupsAnswers, groupAnswers)
			groupAnswers = make(map[rune]bool)
		}
	}

	groupsAnswers = append(groupsAnswers, groupAnswers)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return groupsAnswers
}

func countAnswers(groupsAnswers []map[rune]bool) int {
	var totalAnswers int = 0
	for _, group := range groupsAnswers {
		totalAnswers += len(group)
	}

	return totalAnswers
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("You must supply a file to process.")
	}
	filename := args[0]
	groupsAnswers := processFile(filename)

	fmt.Println("Total answers: ", countAnswers(groupsAnswers))
}
