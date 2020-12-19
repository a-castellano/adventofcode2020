// Álvaro Castellano Vela 2020/12/19
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func processOperation(operation string) int {
	var operationLength int = len(operation)
	var result int = 0
	var position int = 0
	for position < operationLength {
		var operator byte
		var righOperand int
		if position == 0 {
			operator = '+'
		} else {
			operator = operation[position]
			position += 2
		}
		if operation[position] >= '0' && operation[position] <= '9' {
			righOperand = int(operation[position] - '0')
		} else if operation[position] == '(' {
			var openParentheses int = 1
			var openParenthesesPosition int = position
			for openParentheses != 0 {
				position++
				if operation[position] == ')' {
					openParentheses--
				} else if operation[position] == '(' {
					openParentheses++
				}
			}
			righOperand = processOperation(operation[openParenthesesPosition+1 : position])
		}
		if operator == '+' {
			result += righOperand
		} else if operator == '*' {
			result *= righOperand
		}
		position += 2
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
