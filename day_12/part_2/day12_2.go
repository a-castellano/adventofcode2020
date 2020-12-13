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

func recalculateWaypoint(WaypointX int, WaypointY int, instruction rune, steps int) (int, int) {

	var newWaypointX int = 0
	var newWaypointY int = 0

	if instruction == 'R' {
		if steps == 90 {
			newWaypointX = -WaypointY
			newWaypointY = WaypointX
		} else if steps == 180 {
			newWaypointX = -WaypointX
			newWaypointY = -WaypointY
		} else if steps == 270 {
			newWaypointX = WaypointY
			newWaypointY = -WaypointX
		}
	}
	if instruction == 'L' {
		if steps == 90 {
			newWaypointX = WaypointY
			newWaypointY = -WaypointX
		} else if steps == 180 {
			newWaypointX = -WaypointX
			newWaypointY = -WaypointY
		} else if steps == 270 {
			newWaypointX = -WaypointY
			newWaypointY = WaypointX
		}
	}

	return newWaypointX, newWaypointY
}

func processFile(filename string) int {

	var X int = 0
	var Y int = 0

	var WaypointX int = 10
	var WaypointY int = -1

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
			WaypointX, WaypointY = recalculateWaypoint(WaypointX, WaypointY, instruction, steps)
		} else if instruction == 'N' {
			WaypointY -= steps
		} else if instruction == 'E' {
			WaypointX += steps
		} else if instruction == 'W' {
			WaypointX -= steps
		} else if instruction == 'S' {
			WaypointY += steps
		} else if instruction == 'F' {
			X += WaypointX * steps
			Y += WaypointY * steps
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
