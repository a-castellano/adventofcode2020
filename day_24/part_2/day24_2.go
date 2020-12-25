// Álvaro Castellano Vela 2020/12/25
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Tile struct {
	I     int
	J     int
	Color bool
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

func stringfy(i, j int) string {
	var result string
	result = strconv.Itoa(i) + "|" + strconv.Itoa(j)
	return result
}

func flipTiles(directions [][]string) map[string]Tile {

	blackTiles := make(map[string]Tile)

	tiles := make([][]bool, 1000)
	// Flase -> White
	// True  -> Trve Black
	for x := 0; x < len(tiles); x++ {
		line := make([]bool, 1000)
		tiles[x] = line
	}

	for _, direction := range directions {
		i, j := calculatePosition(500, 500, direction)
		if !tiles[i][j] {
			tiles[i][j] = true
			tile := Tile{i, j, true}
			blackTiles[stringfy(i, j)] = tile
		} else {
			tiles[i][j] = false
			delete(blackTiles, stringfy(i, j))
		}
	}

	return blackTiles
}

func calculateAdjecent(tile Tile) []string {

	var i, j int
	var adjacentStrings []string

	direction := make([]string, 1)

	direction[0] = "e"
	i, j = calculatePosition(tile.I, tile.J, direction)
	adjacentStrings = append(adjacentStrings, stringfy(i, j))

	direction[0] = "w"
	i, j = calculatePosition(tile.I, tile.J, direction)
	adjacentStrings = append(adjacentStrings, stringfy(i, j))

	direction[0] = "ne"
	i, j = calculatePosition(tile.I, tile.J, direction)
	adjacentStrings = append(adjacentStrings, stringfy(i, j))

	direction[0] = "nw"
	i, j = calculatePosition(tile.I, tile.J, direction)
	adjacentStrings = append(adjacentStrings, stringfy(i, j))

	direction[0] = "se"
	i, j = calculatePosition(tile.I, tile.J, direction)
	adjacentStrings = append(adjacentStrings, stringfy(i, j))

	direction[0] = "sw"
	i, j = calculatePosition(tile.I, tile.J, direction)
	adjacentStrings = append(adjacentStrings, stringfy(i, j))

	return adjacentStrings
}

func stringToTile(position string, color bool) Tile {
	var tile Tile

	splitedPosition := strings.Split(position, "|")
	tile.I, _ = strconv.Atoi(splitedPosition[0])
	tile.J, _ = strconv.Atoi(splitedPosition[1])
	tile.Color = color

	return tile
}

func applyDays(blackTiles map[string]Tile, days int) map[string]Tile {

	var currentBlackTiles map[string]Tile = blackTiles

	for days > 0 {
		newBlackTiles := make(map[string]Tile)
		whiteTiles := make(map[string]Tile)
		// generate inmediate neighbors
		for _, tile := range currentBlackTiles {
			adjacentTiles := calculateAdjecent(tile)
			// Create current white tiles
			for _, adjacentTile := range adjacentTiles {
				if _, ok := currentBlackTiles[adjacentTile]; !ok {
					if _, ok := whiteTiles[adjacentTile]; !ok {
						whiteTiles[adjacentTile] = stringToTile(adjacentTile, false) //white
					}
				}
			}
		}
		// generate whiteTiles neighbors
		for _, tile := range whiteTiles {
			adjacentTiles := calculateAdjecent(tile)
			// Create current white tiles
			for _, adjacentTile := range adjacentTiles {
				if _, ok := currentBlackTiles[adjacentTile]; !ok {
					if _, ok := whiteTiles[adjacentTile]; !ok {
						whiteTiles[adjacentTile] = stringToTile(adjacentTile, false) //white
					}
				}
			}
		}

		//caclulate newBlackTiles from currentBlackTiles
		for stringTile, tile := range currentBlackTiles {
			adjacentTiles := calculateAdjecent(tile)
			var adjacentBlack int = 0
			for _, adjacentTile := range adjacentTiles {
				if _, ok := currentBlackTiles[adjacentTile]; ok {
					adjacentBlack++
				}
			}
			if adjacentBlack == 1 || adjacentBlack == 2 { // Black is maintained
				newBlackTiles[stringTile] = tile
			}
		}
		//expand newBlackTiles from whiteTiles
		for stringTile, tile := range whiteTiles {
			adjacentTiles := calculateAdjecent(tile)
			var adjacentBlack int = 0
			for _, adjacentTile := range adjacentTiles {
				if _, ok := currentBlackTiles[adjacentTile]; ok {
					adjacentBlack++
				}
			}
			if adjacentBlack == 2 { // Black is maintained
				newBlackTiles[stringTile] = tile
			}
		}

		currentBlackTiles = newBlackTiles
		days--
	}

	return currentBlackTiles
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("You must supply a file to process.")
	}
	filename := args[0]

	directions := processFile(filename)

	currentBlackTiles := flipTiles(directions)
	finalBlackTile := applyDays(currentBlackTiles, 100)

	fmt.Println(len(finalBlackTile))
}
