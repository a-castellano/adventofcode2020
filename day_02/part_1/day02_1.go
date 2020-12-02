// Álvaro Castellano Vela 2020/12/01
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func processFile(filename string) int {

	var validPasswords int = 0

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	passwordRe := regexp.MustCompile("^([[:digit:]]+)-([[:digit:]]+) ([a-z]): ([a-z]+)$")

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		var ocurrences int = 0
		match := passwordRe.FindAllStringSubmatch(scanner.Text(), -1)
		min, _ := strconv.Atoi(match[0][1])
		max, _ := strconv.Atoi(match[0][2])
		pattern := []rune(match[0][3])[0]
		password := match[0][4]
		for _, char := range password {
			if char == pattern {
				ocurrences++
			}
		}
		if ocurrences >= min && ocurrences <= max {
			validPasswords++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return validPasswords
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("You must supply a file to process.")
	}
	filename := args[0]
	validPasswords := processFile(filename)
	fmt.Printf("Valid Passwords: %d\n", validPasswords)
}
