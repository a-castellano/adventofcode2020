// Álvaro Castellano Vela 2020/12/04
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func processFile(filename string) []string {

	var candidates []string
	var auxLine string

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			auxLine += " " + line
		} else {
			candidates = append(candidates, auxLine)
			auxLine = ""
		}
	}
	candidates = append(candidates, auxLine)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return candidates
}

func countValidDocs(candidates []string) int {
	var valid_docs int = 0
	required_items := map[string]bool{"byr": true, "iyr": true, "eyr": true, "hgt": true, "hcl": true, "ecl": true, "pid": true}
	for _, candidate := range candidates {
		var valid_items int = 0
		parts := strings.Split(candidate, " ")[1:]
		for _, part := range parts {
			key := strings.Split(part, ":")[0]
			if _, ok := required_items[key]; ok {
				valid_items++
			}
		}

		if valid_items == 7 {
			valid_docs++
		}
	}

	return valid_docs
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("You must supply a file to process.")
	}
	filename := args[0]
	candidates := processFile(filename)

	valid_docs := countValidDocs(candidates)

	fmt.Println("Valid docs: ", valid_docs)
}
