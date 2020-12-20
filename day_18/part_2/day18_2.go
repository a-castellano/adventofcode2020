// Álvaro Castellano Vela 2020/12/19
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func processOperation(operation string) int {
	var result int = 0
	//look for parentheses
	var hasParentheses bool = true

	for hasParentheses {
		var operationLength int = len(operation)
		var position int = 0
		var foundParentheses bool = false
		var initialParentheses int
		var finalParentheses int
		for position < operationLength {
			if operation[position] == '(' {
				initialParentheses = position
				foundParentheses = true
				var openParentheses int = 1
				for openParentheses != 0 {
					position++
					if operation[position] == ')' {
						openParentheses--
					} else if operation[position] == '(' {
						openParentheses++
					}
				}
				finalParentheses = position
				break
			}
			position++
		}
		if !foundParentheses {
			hasParentheses = false
		} else {
			parenthesesResult := (processOperation(operation[initialParentheses+1 : finalParentheses]))
			newOperation := fmt.Sprintf("%s%d%s", operation[:initialParentheses], parenthesesResult, operation[finalParentheses+1:])
			operation = newOperation
		}
	}

	items := strings.Split(operation, " ")
	// look for +

	var hasPlus bool = true

	for hasPlus {
		var foundPlus bool = false
		for postition, item := range items {
			if item == "+" {
				foundPlus = true
				left, _ := strconv.Atoi(items[postition-1])
				right, _ := strconv.Atoi(items[postition+1])
				newValue := left + right
				newItems := append(items[:postition-1], strconv.Itoa(newValue))
				newItems = append(newItems, items[postition+2:]...)
				items = newItems
				break
			}
		}
		if !foundPlus {
			hasPlus = false
		}
	}

	// Only * left in this operation
	if len(items) > 1 {
		result = 1
		for _, item := range items {
			if item != "*" {
				itemNumber, _ := strconv.Atoi(item)
				result *= itemNumber
			}
		}
	} else {
		result, _ = strconv.Atoi(items[0])
	}
	return result
}

func processFile(filename string) int {

	var result int = 0

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// read map
	for scanner.Scan() {
		operation := scanner.Text()
		result += processOperation(operation)
	}

	return result
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("You must supply a file to process.")
	}
	filename := args[0]

	result := processFile(filename)
	fmt.Println("Result:", result)
}
