// Álvaro Castellano Vela 2020/12/20
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

func expandRule(rules map[int]string, expandedRules map[int]bool, ruleId int) string {

	fmt.Println("Expading", ruleId, rules[ruleId])

	if _, ok := expandedRules[ruleId]; !ok {
		var stringRule string = rules[ruleId]
		var orFound bool = false
		if stringRule[0] == '"' {
			rules[ruleId] = fmt.Sprintf("%c", stringRule[1])
			expandedRules[ruleId] = true
		} else {
			splitedString := strings.Split(rules[ruleId], " ")
			var newString string
			fmt.Println("Expading", ruleId, "newString", newString)
			for _, id := range splitedString {
				// Check if it is an actual id
				if id[0] >= '0' && id[0] <= '9' {

					//expand Rule
					idInt, _ := strconv.Atoi(id)
					newString = fmt.Sprintf("%s%s", newString, expandRule(rules, expandedRules, idInt))
					fmt.Println("Expading", ruleId, "newString", newString)
				} else if id[0] == '|' {
					orFound = true
					newString = fmt.Sprintf("(%s|", newString)
				}
			}
			if orFound {
				newString = fmt.Sprintf("%s)", newString)
			}
			rules[ruleId] = newString
			expandedRules[ruleId] = true
		}
	}
	return rules[ruleId]
}

func processFile(filename string) (map[int]string, []string) {

	rules := make(map[int]string)
	expandedRules := make(map[int]bool)
	candidates := make([]string, 0)

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	ruleRe := regexp.MustCompile(`^([0-9]+): (.*)$`)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// read rules
	for scanner.Scan() {
		ruleString := scanner.Text()
		if ruleString == "" {
			break
		} else {
			match := ruleRe.FindAllStringSubmatch(ruleString, -1)
			ruleId, _ := strconv.Atoi(match[0][1])
			rules[ruleId] = match[0][2]
		}
	}

	// Read candidates
	for scanner.Scan() {
		candidates = append(candidates, scanner.Text())
	}

	expandRule(rules, expandedRules, 0)
	return rules, candidates
}

func validateCandidates(rules map[int]string, candidates []string, ruleId int) int {
	var validCandidates int = 0

	regex := fmt.Sprintf("^%s$", rules[ruleId])
	ruleRe := regexp.MustCompile(regex)
	for _, candidate := range candidates {
		match := ruleRe.FindAllStringSubmatch(candidate, -1)
		if len(match) == 1 {
			validCandidates++
		}
	}
	return validCandidates
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("You must supply a file to process.")
	}
	filename := args[0]

	rules, candidates := processFile(filename)

	fmt.Println("Rules:", rules)
	fmt.Println("Rule 0:", rules[0])
	fmt.Println("Valid candidates:", validateCandidates(rules, candidates, 0))
}
