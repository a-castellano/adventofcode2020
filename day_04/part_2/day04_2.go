// Álvaro Castellano Vela 2020/12/05
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

	// Required regular expressions
	hgtRe := regexp.MustCompile("^([[:digit:]]+)(cm|in)")
	hclRe := regexp.MustCompile("^#[0-9a-z]{6}$")
	eclRe := regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$")
	pidRe := regexp.MustCompile("^[0-9]{9}$")

	for _, candidate := range candidates {
		var valid_items int = 0
		parts := strings.Split(candidate, " ")[1:]
		var valid bool = true
		for _, part := range parts {
			key := strings.Split(part, ":")[0]
			value := strings.Split(part, ":")[1]
			if _, ok := required_items[key]; ok {
				valid_items++
				if key == "byr" {
					year, _ := strconv.Atoi(value)
					if year < 1920 || year > 2002 {
						valid = false
					}
				} else if key == "iyr" {
					year, _ := strconv.Atoi(value)
					if year < 2010 || year > 2020 {
						valid = false
					}
				} else if key == "eyr" {
					year, _ := strconv.Atoi(value)
					if year < 2020 || year > 2030 {
						valid = false
					}
				} else if key == "hgt" {
					match := hgtRe.FindAllStringSubmatch(value, -1)
					if len(match) != 1 {
						valid = false
					} else {
						unit := match[0][2]
						unit_value, _ := strconv.Atoi(match[0][1])
						if unit == "cm" {
							if unit_value < 150 || unit_value > 193 {
								valid = false
							}
						} else { // unit == "in"
							if unit_value < 59 || unit_value > 76 {
								valid = false
							}
						}
					}
				} else if key == "hcl" {
					match := hclRe.FindAllStringSubmatch(value, -1)
					if len(match) != 1 {
						valid = false
					}
				} else if key == "ecl" {
					match := eclRe.FindAllStringSubmatch(value, -1)
					if len(match) != 1 {
						valid = false
					}
				} else if key == "pid" {
					match := pidRe.FindAllStringSubmatch(value, -1)
					if len(match) != 1 {
						valid = false
					}
				}

			}
		}

		if valid_items == 7 && valid == true {
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
