// Álvaro Castellano Vela 2020/12/20
package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Borders struct {
	N string
	S string
	E string
	W string
}

type Tile struct {
	ID              int
	Orientation     int
	N               int
	S               int
	E               int
	W               int
	Rotations       [][][]rune
	RotationBorders []Borders
}

func getBorders(image [][]rune) Borders {

	var borders Borders
	var e, w []rune

	borders.N = string(image[0])
	borders.S = string(image[len(image)-1])

	for i := 0; i < len(image); i++ {
		w = append(w, image[i][0])
	}
	for i := 0; i < len(image); i++ {
		e = append(e, image[i][len(image[i])-1])
	}

	borders.E = string(e)
	borders.W = string(w)
	return borders
}

func rotate(image [][]rune, heigh, width int) [][]rune {
	var rotated [][]rune

	for i := 0; i < heigh; i++ {
		row := make([]rune, width)
		rotated = append(rotated, row)
	}

	for i := 0; i < heigh; i++ {
		for j := 0; j < width; j++ {
			rotated[j][heigh-1-i] = image[i][j]
		}
	}

	return rotated
}

func flip(image [][]rune, heigh, width int) [][]rune {
	var fliped [][]rune

	for i := 0; i < heigh; i++ {
		row := make([]rune, width)
		fliped = append(fliped, row)
	}

	for i := 0; i < heigh; i++ {
		for j := 0; j < width; j++ {
			fliped[i][width-1-j] = image[i][j]
		}
	}

	return fliped
}

func processFile(filename string) (map[int]Tile, [][]int) {

	tiles := make(map[int]Tile)
	var heigh int = 0
	var width int = 0
	var tilesPositions [][]int

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	tileIDRe := regexp.MustCompile(`^Tile ([0-9]+):$`)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// read Tiles
	for scanner.Scan() {
		tileIDString := scanner.Text()
		match := tileIDRe.FindAllStringSubmatch(tileIDString, -1)
		tileId, _ := strconv.Atoi(match[0][1])
		tile := Tile{tileId, 0, -1, -1, -1, -1, make([][][]rune, 8), make([]Borders, 8)}
		for scanner.Scan() {
			tileLine := scanner.Text()
			if tileLine != "" {
				var tileSlice []rune
				for _, value := range tileLine {
					tileSlice = append(tileSlice, value)
				}
				tile.Rotations[0] = append(tile.Rotations[0], tileSlice)
			} else {
				break
			}
		}
		tiles[tileId] = tile
		heigh = len(tiles[tileId].Rotations[0])
		width = len(tiles[tileId].Rotations[0][0])
	}

	for _, tile := range tiles {
		tile.Rotations[1] = rotate(tile.Rotations[0], heigh, width)
		tile.Rotations[2] = rotate(tile.Rotations[1], heigh, width)
		tile.Rotations[3] = rotate(tile.Rotations[2], heigh, width)
		tile.Rotations[4] = flip(tile.Rotations[0], heigh, width)
		tile.Rotations[5] = flip(tile.Rotations[1], heigh, width)
		tile.Rotations[6] = flip(tile.Rotations[2], heigh, width)
		tile.Rotations[7] = flip(tile.Rotations[3], heigh, width)
		tile.RotationBorders[0] = getBorders(tile.Rotations[0])
		tile.RotationBorders[1] = getBorders(tile.Rotations[1])
		tile.RotationBorders[2] = getBorders(tile.Rotations[2])
		tile.RotationBorders[3] = getBorders(tile.Rotations[3])
		tile.RotationBorders[4] = getBorders(tile.Rotations[4])
		tile.RotationBorders[5] = getBorders(tile.Rotations[5])
		tile.RotationBorders[6] = getBorders(tile.Rotations[6])
		tile.RotationBorders[7] = getBorders(tile.Rotations[7])
	}

	tileLengh := int(math.Sqrt(float64(len(tiles))))

	for i := 0; i < tileLengh; i++ {
		row := make([]int, tileLengh)
		tilesPositions = append(tilesPositions, row)
	}

	return tiles, tilesPositions
}

func showTile(rotation [][]rune) {
	for _, row := range rotation {
		for _, column := range row {
			fmt.Printf("%c", column)
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func showTiles(tiles map[int]Tile) {
	for _, tile := range tiles {
		fmt.Println("Tile: ", tile.ID)
		for _, rotation := range tile.Rotations {
			showTile(rotation)
		}
	}
}

func findEdges(tiles map[int]Tile, tilesPositions [][]int) int {

	edgeCandidates := make(map[int][]int)
	var result int = 1

	//find candidates
	for tileId, tile := range tiles {
		if _, ok := edgeCandidates[tileId]; !ok {
			for rotation, _ := range tile.Rotations {
				var matches int = 0
				for candidateTileId, candidateTile := range tiles {
					if tileId != candidateTileId {
						for candidateRotation, _ := range candidateTile.Rotations {
							if tile.RotationBorders[rotation].N == candidateTile.RotationBorders[candidateRotation].S || tile.RotationBorders[rotation].W == candidateTile.RotationBorders[candidateRotation].E {
								matches++
							}
						}
					}
				}
				if matches == 0 {
					if _, ok := edgeCandidates[tileId]; !ok {
						edgeCandidates[tileId] = make([]int, 0)
					}
					edgeCandidates[tileId] = append(edgeCandidates[tileId], rotation)
				}
			}
		}
	}

	for tileId, _ := range edgeCandidates {
		result *= tileId
	}
	return result
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("You must supply a file to process.")
	}
	filename := args[0]

	tiles, tilesPositions := processFile(filename)

	result := findEdges(tiles, tilesPositions)
	fmt.Println("Result:", result)
}
