// Álvaro Castellano Vela 2020/12/25
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Tyle struct {
	X int
	Y int
}

func processFile(filename string) [][]string {

	var directions [][]string

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		directionsString := scanner.Text()
		var direction []string
		for i := 0; i < len(directionsString); i++ {
			if directionsString[i] == 'e' {
				direction = append(direction, "e")
			} else if directionsString[i] == 'w' {
				direction = append(direction, "w")
			} else if directionsString[i] == 'n' {
				if directionsString[i+1] == 'e' {
					direction = append(direction, "ne")
				} else {
					direction = append(direction, "nw")
				}
				i++
			} else if directionsString[i] == 's' {
				if directionsString[i+1] == 'e' {
					direction = append(direction, "se")
				} else {
					direction = append(direction, "sw")
				}
				i++
			}
		}
		directions = append(directions, direction)
	}

	return directions
}

func calculatePosition(startX int, startY int, direction []string) (int, int) {
	var x int = startX
	var y int = startY
	for _, step := range direction {
		if step == "e" {
			y++
		} else if step == "w" {
			y--
		} else if step == "ne" {
			x--
			if x%2 == 0 {
				y++
			}
		} else if step == "nw" {
			x--
			if x%2 != 0 {
				y--
			}
		} else if step == "se" {
			x++
			if x%2 == 0 {
				y++
			}
		} else if step == "sw" {
			x++
			if x%2 != 0 {
				y--
			}
		}
	}
	return x, y
}

func flipTiles(directions [][]string) int {

	var blackTiles int = 0

	tiles := make([][]bool, 1000)
	// Flase -> White
	// True  -> Trve Black
	for i := 0; i < len(tiles); i++ {
		line := make([]bool, 1000)
		tiles[i] = line
	}

	for _, direction := range directions {
		i, j := calculatePosition(500, 500, direction)
		if !tiles[i][j] {
			tiles[i][j] = true
			blackTiles++
		} else {
			tiles[i][j] = false
			blackTiles--
		}
	}

	return blackTiles
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("You must supply a file to process.")
	}
	filename := args[0]

	directions := processFile(filename)

	fmt.Println("Black tiles:", flipTiles(directions))
}
