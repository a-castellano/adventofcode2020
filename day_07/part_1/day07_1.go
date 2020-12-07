// Álvaro Castellano Vela 2020/12/07
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	//	"strings"
)

type Bag struct {
	Color    string
	Contains map[string]int
}

func processFile(filename string) map[string]Bag {

	bagRules := make(map[string]Bag)

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	bagRule := regexp.MustCompile("^([a-z]+) ([a-z]+) bags contain (no other bags|.+).$")
	bahPropertiesRule := regexp.MustCompile("([0-9]+) ([a-z]+) ([a-z]+) bag")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		match := bagRule.FindAllStringSubmatch(scanner.Text(), -1)
		bagColor := match[0][1] + " " + match[0][2]

		if _, ok := bagRules[bagColor]; !ok {
			newBag := Bag{bagColor, make(map[string]int)}
			bagRules[bagColor] = newBag
		}

		bagPropertiesString := match[0][3]
		if bagPropertiesString != "no other bags" {
			matchProperties := bahPropertiesRule.FindAllStringSubmatch(bagPropertiesString, -1)
			for _, property := range matchProperties {
				number, _ := strconv.Atoi(property[1])
				color := property[2] + " " + property[3]
				bagRules[bagColor].Contains[color] = number
				// Get all bags, some of them do not have properties line
				if _, ok := bagRules[color]; !ok {
					childBag := Bag{color, make(map[string]int)}
					bagRules[color] = childBag
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return bagRules
}

func canContain(bagRules map[string]Bag, canContainMap map[string]bool, bagColor string, requiredBag string) {

	if len(bagRules[bagColor].Contains) == 0 {
		canContainMap[bagColor] = false
	} else {
		found := false
		for color, _ := range bagRules[bagColor].Contains {
			if color == requiredBag {
				canContainMap[bagColor] = true
				found = true
			}
			if _, ok := canContainMap[color]; !ok {
				canContain(bagRules, canContainMap, color, requiredBag)
			}
			if canContainMap[color] {
				//sub bag can contain so bagColor too
				canContainMap[bagColor] = true
				found = true
			}
		}
		if !found {
			canContainMap[bagColor] = false
		}
	}
}

func howManyCanContain(bagRules map[string]Bag, requiredBag string) int {

	canContainMap := make(map[string]bool)
	var containCounter int = 0

	for bagColor, bag := range bagRules {
		if _, ok := bag.Contains[requiredBag]; ok {
			// bag can contain our required bag, incremnet and stop searching
			canContainMap[bagColor] = true
			containCounter++
		} else {
			// one of the sub bags may contain our bag, check it
			for color, _ := range bag.Contains {
				if _, ok := canContainMap[color]; !ok {
					canContain(bagRules, canContainMap, color, requiredBag)
				}
				if canContainMap[color] {
					canContainMap[bagColor] = true
					containCounter++
					break
				}
			}
		}
	}
	return containCounter
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("You must supply a file to process.")
	}
	filename := args[0]
	bagRules := processFile(filename)
	result := howManyCanContain(bagRules, "shiny gold")

	fmt.Println("There are", result, "bags that can contain a 'shiny gold' bag.")
}
