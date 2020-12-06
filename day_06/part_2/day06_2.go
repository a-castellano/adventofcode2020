// Álvaro Castellano Vela 2020/12/06
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func processFile(filename string) ([]map[rune]int, []int) {

	groupsAnswers := make([]map[rune]int, 0)
	var groupsMembers []int

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	groupAnswers := make(map[rune]int)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var groupMembers = 0

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			groupMembers++
			for _, answer := range line {
				if _, ok := groupAnswers[answer]; ok {
					groupAnswers[answer]++
				} else {
					groupAnswers[answer] = 1
				}
			}
		} else {
			groupsAnswers = append(groupsAnswers, groupAnswers)
			groupsMembers = append(groupsMembers, groupMembers)
			groupAnswers = make(map[rune]int)
			groupMembers = 0
		}
	}

	groupsAnswers = append(groupsAnswers, groupAnswers)
	groupsMembers = append(groupsMembers, groupMembers)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return groupsAnswers, groupsMembers
}

func countAnswers(groupsAnswers []map[rune]int, groupsMembers []int) int {
	var totalAnswers int = 0
	for index, group := range groupsAnswers {
		for _, answers := range group {
			if groupsMembers[index] == answers {
				totalAnswers++
			}
		}
	}

	return totalAnswers
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("You must supply a file to process.")
	}
	filename := args[0]
	groupsAnswers, groupsMembers := processFile(filename)

	fmt.Println("Total answers: ", countAnswers(groupsAnswers, groupsMembers))
}
