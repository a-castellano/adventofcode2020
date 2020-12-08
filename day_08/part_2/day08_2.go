// Álvaro Castellano Vela 2020/12/08
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

type Instruction struct {
	Name   string
	Offset int
}

func processFile(filename string) ([]Instruction, []int) {

	intructions := make([]Instruction, 0)
	var jumpsOrNops []int
	var counter int = 0

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	instructionRule := regexp.MustCompile("^(acc|jmp|nop) (.)([0-9]+)$")

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		match := instructionRule.FindAllStringSubmatch(scanner.Text(), -1)
		var instruction Instruction
		instruction.Name = match[0][1]
		instruction.Offset, _ = strconv.Atoi(match[0][3])
		if match[0][2] == "-" {
			instruction.Offset = 0 - instruction.Offset
		}
		intructions = append(intructions, instruction)
		if instruction.Name == "jmp" || instruction.Name == "nop" {
			jumpsOrNops = append(jumpsOrNops, counter)
		}
		counter++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return intructions, jumpsOrNops
}

func findLoop(instructions []Instruction) (int, bool) {
	var accumulator int = 0
	var counter int = 0
	var loop bool = false
	var instructionsLenght int = len(instructions)

	executesInstructions := make(map[int]bool)
	executesInstructions[counter] = true

	for {
		instruction := instructions[counter]
		if instruction.Name == "acc" {
			accumulator += instruction.Offset
			counter++
		} else if instruction.Name == "nop" {
			counter++
		} else if instruction.Name == "jmp" {
			counter += instruction.Offset
		}
		if _, ok := executesInstructions[counter]; !ok {
			executesInstructions[counter] = true
			if counter == instructionsLenght { // end of our porgram
				break
			}
		} else {
			loop = true
			break
		}
	}

	return accumulator, loop
}

func findWrongInstruction(instructions []Instruction, jumpsOrNops []int) int {
	var accumulator int
	var instructionsLenght int = len(instructions)

	for _, instruction := range jumpsOrNops {
		copiedInstructions := make([]Instruction, instructionsLenght)
		copy(copiedInstructions, instructions)
		if copiedInstructions[instruction].Name == "jmp" {
			copiedInstructions[instruction].Name = "nop"
		} else {
			copiedInstructions[instruction].Name = "jmp"
		}
		acc, loop := findLoop(copiedInstructions)
		if !loop {
			accumulator = acc
			break
		}
	}

	return accumulator
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("You must supply a file to process.")
	}
	filename := args[0]

	intructions, jumpsOrNops := processFile(filename)
	accumulator := findWrongInstruction(intructions, jumpsOrNops)

	fmt.Println("Accumulator value:", accumulator)
}
