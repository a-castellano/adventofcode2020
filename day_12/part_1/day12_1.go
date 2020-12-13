// Álvaro Castellano Vela 2020/12/13
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func abs(n int) int {

	if n < 0 {
		return -n
	} else {
		return n
	}
}

func calculateNewDirection(currentDirection string, instruction rune, steps int) string {
	var newDirection string
	if instruction == 'R' {
		if currentDirection == "east" {
			if steps == 90 {
				newDirection = "south"
			} else if steps == 180 {
				newDirection = "west"
			} else if steps == 270 {
				newDirection = "north"
			}
		}
		if currentDirection == "south" {
			if steps == 90 {
				newDirection = "west"
			} else if steps == 180 {
				newDirection = "north"
			} else if steps == 270 {
				newDirection = "east"
			}
		}
		if currentDirection == "west" {
			if steps == 90 {
				newDirection = "north"
			} else if steps == 180 {
				newDirection = "east"
			} else if steps == 270 {
				newDirection = "south"
			}
		}
		if currentDirection == "north" {
			if steps == 90 {
				newDirection = "east"
			} else if steps == 180 {
				newDirection = "south"
			} else if steps == 270 {
				newDirection = "west"
			}
		}
	}
	if instruction == 'L' {
		if currentDirection == "east" {
			if steps == 90 {
				newDirection = "north"
			} else if steps == 180 {
				newDirection = "west"
			} else if steps == 270 {
				newDirection = "south"
			}
		}
		if currentDirection == "north" {
			if steps == 90 {
				newDirection = "west"
			} else if steps == 180 {
				newDirection = "south"
			} else if steps == 270 {
				newDirection = "east"
			}
		}
		if currentDirection == "west" {
			if steps == 90 {
				newDirection = "south"
			} else if steps == 180 {
				newDirection = "east"
			} else if steps == 270 {
				newDirection = "north"
			}
		}
		if currentDirection == "south" {
			if steps == 90 {
				newDirection = "east"
			} else if steps == 180 {
				newDirection = "north"
			} else if steps == 270 {
				newDirection = "west"
			}
		}
	}

	return newDirection
}

func processFile(filename string) int {

	var currentDirection string = "east"
	var X int
	var Y int

	instructionRule := regexp.MustCompile("^(N|E|S|W|R|L|F)([0-9]+)$")

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		match := instructionRule.FindAllStringSubmatch(scanner.Text(), -1)
		instruction := []rune(match[0][1])[0]
		steps, _ := strconv.Atoi(match[0][2])
		if instruction == 'L' || instruction == 'R' {
			currentDirection = calculateNewDirection(currentDirection, instruction, steps)
		} else if instruction == 'N' {
			Y -= steps
		} else if instruction == 'E' {
			X += steps
		} else if instruction == 'W' {
			X -= steps
		} else if instruction == 'S' {
			Y += steps
		} else if instruction == 'F' {
			if currentDirection == "north" {
				Y -= steps
			} else if currentDirection == "east" {
				X += steps
			} else if currentDirection == "west" {
				X -= steps
			} else if currentDirection == "south" {
				Y += steps
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return abs(X) + abs(Y)
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("You must supply a file to process.")
	}
	filename := args[0]

	manhattanDistance := processFile(filename)

	fmt.Println("Manhattan distance:", manhattanDistance)
}
