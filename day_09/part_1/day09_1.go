// Álvaro Castellano Vela 2020/12/09
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func processFile(filename string) []int {

	var numbers []int

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return numbers
}

func checkValidNumber(numbers []int, firstPosition int, lastPosition int, candidate int) bool {

	var valid bool = false

	for i := firstPosition; i <= lastPosition; i++ {
		for j := firstPosition; j <= lastPosition; j++ {
			if i != j {
				if numbers[i]+numbers[j] == candidate {
					valid = true
					break
				}
			}
		}
	}

	return valid

}

func findNumber(numbers []int, preamble int) int {

	var number int = -1
	var firstPosition int = 0
	var lastPosition int = preamble - 1

	for i := lastPosition + 1; i < len(numbers); i++ {

		candidate := numbers[i]

		if !checkValidNumber(numbers, firstPosition, lastPosition, candidate) {

			number = candidate
			break

		}

		firstPosition++
		lastPosition++

	}

	return number
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatal("You must supply a file to process, and preamble number.")
	}
	filename := args[0]
	preamble, _ := strconv.Atoi(args[1])

	numbers := processFile(filename)
	result := findNumber(numbers, preamble)
	fmt.Println("Result:", result)
}
