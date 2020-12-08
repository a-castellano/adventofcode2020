// Álvaro Castellano Vela 2020/12/07
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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

func countInside(bagRules map[string]Bag, countMap map[string]int, requiredBag string) int {
	var counter int = 0

	if len(bagRules[requiredBag].Contains) == 0 {
		return 0
	} else {
		for color, quantity := range bagRules[requiredBag].Contains {
			if _, ok := countMap[color]; !ok {
				countMap[color] = countInside(bagRules, countMap, color)
			}
			counter += quantity*countMap[color] + quantity
		}
		return counter
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("You must supply a file to process.")
	}
	filename := args[0]
	bagRules := processFile(filename)
	countMap := make(map[string]int)
	result := countInside(bagRules, countMap, "shiny gold")
	fmt.Println("A single shiny gold bag must contain", result, "other bags.")
}
