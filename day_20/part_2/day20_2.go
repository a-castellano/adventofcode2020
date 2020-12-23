// Álvaro Castellano Vela 2020/12/23
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

func processFile(filename string) map[int]Tile {

	tiles := make(map[int]Tile)
	var heigh int = 0
	var width int = 0

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

	return tiles
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
		fmt.Println(tile.RotationBorders)
	}
}

func findEdges(tiles map[int]Tile) map[int][]int {

	edgeCandidates := make(map[int][]int)

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
	return edgeCandidates
}

// There can be 8 different combinations, all valid, generate only one
func createImage(tiles map[int]Tile, edges map[int][]int) [][]rune {

	var image [][]rune

	var startTile int
	var startRotation int
	usedTiles := make(map[int]bool)
	var imageLength int

	tilesPositions := make([][]int, 0)
	tilesRotations := make([][]int, 0)

	tileLengh := int(math.Sqrt(float64(len(tiles))))

	for i := 0; i < tileLengh; i++ {
		row := make([]int, tileLengh)
		tilesPositions = append(tilesPositions, row)
	}
	for i := 0; i < tileLengh; i++ {
		row := make([]int, tileLengh)
		tilesRotations = append(tilesRotations, row)
	}

	for tileId, rotations := range edges {
		startTile = tileId
		startRotation = rotations[0]
		break
	}
	tilesPositions[0][0] = startTile
	tilesRotations[0][0] = startRotation
	usedTiles[startTile] = true

	for i := 0; i < tileLengh; i++ {
		for j := 0; j < tileLengh; j++ {
			//var tileFound bool = false
			if tilesPositions[i][j] == 0 { //edges will be marked as ocuppied
				if i == 0 { // Uppder border
					var validCandidates int = 0
					for tileId, tile := range tiles {
						if _, ok := usedTiles[tileId]; !ok {
							for rotation, _ := range tile.Rotations {
								if tile.RotationBorders[rotation].W == tiles[tilesPositions[i][j-1]].RotationBorders[tilesRotations[i][j-1]].E {
									tilesPositions[i][j] = tileId
									tilesRotations[i][j] = rotation
									usedTiles[tileId] = true
									validCandidates++
								}
							}
						}
					}
				} else {
					if j == 0 {
						var validCandidates int = 0
						for tileId, tile := range tiles {
							if _, ok := usedTiles[tileId]; !ok {
								for rotation, _ := range tile.Rotations {
									if tile.RotationBorders[rotation].N == tiles[tilesPositions[i-1][j]].RotationBorders[tilesRotations[i-1][j]].S {
										tilesPositions[i][j] = tileId
										tilesRotations[i][j] = rotation
										usedTiles[tileId] = true
										validCandidates++
									}
								}
							}
						}
					} else { // j>0 compare W and N
						var validCandidates int = 0
						for tileId, tile := range tiles {
							if _, ok := usedTiles[tileId]; !ok {
								for rotation, _ := range tile.Rotations {
									if tile.RotationBorders[rotation].N == tiles[tilesPositions[i-1][j]].RotationBorders[tilesRotations[i-1][j]].S && tile.RotationBorders[rotation].W == tiles[tilesPositions[i][j-1]].RotationBorders[tilesRotations[i][j-1]].E {
										tilesPositions[i][j] = tileId
										tilesRotations[i][j] = rotation
										usedTiles[tileId] = true
										validCandidates++
									}
								}
							}
						}
					}
				}
			}
		}
	}

	// Create empty image
	imageLength = len(tiles[tilesPositions[0][0]].Rotations[0][0])
	for i := 0; i < tileLengh; i++ {
		for j := 0; j < imageLength-2; j++ {
			row := make([]rune, (imageLength-2)*tileLengh)
			image = append(image, row)
		}
	}

	for i := 0; i < tileLengh; i++ {
		for j := 0; j < tileLengh; j++ {
			subImageWithBorders := tiles[tilesPositions[i][j]].Rotations[tilesRotations[i][j]]
			for k := 1; k < imageLength-1; k++ {
				for l := 1; l < imageLength-1; l++ {
					image[i*(imageLength-2)+k-1][j*(imageLength-2)+l-1] = subImageWithBorders[k][l]
				}
			}
		}
	}
	return image
}

func countMonsters(image [][]rune, monster [][]rune, monsterPounds int) int {
	var poundsNoMonster int = 0

	images := make([][][]rune, 8)
	images[0] = image
	images[1] = rotate(images[0], len(image), len(image[0]))
	images[2] = rotate(images[1], len(image), len(image[0]))
	images[3] = rotate(images[2], len(image), len(image[0]))
	images[4] = flip(images[3], len(image), len(image[0]))
	images[5] = flip(images[4], len(image), len(image[0]))
	images[6] = flip(images[5], len(image), len(image[0]))
	images[7] = flip(images[6], len(image), len(image[0]))

	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[0]); j++ {
			if image[i][j] == '#' {
				poundsNoMonster++
			}
		}
	}

	for _, iamgeCandidate := range images {
		for i := 0; i < len(image)-len(monster); i++ {
			for j := 0; j < len(image[0])-len(monster[0]); j++ {
				var poundsLeft int = monsterPounds
				for k := 0; k < len(monster); k++ {
					for l := 0; l < len(monster[0]); l++ {
						if monster[k][l] == iamgeCandidate[i+k][j+l] { //#
							poundsLeft--
						}
					}
				}
				if poundsLeft == 0 {
					poundsNoMonster -= monsterPounds
				}
			}
		}
	}

	return poundsNoMonster
}

func main() {

	monster := [][]rune{
		{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#', ' '},
		{'#', ' ', ' ', ' ', ' ', '#', '#', ' ', ' ', ' ', ' ', '#', '#', ' ', ' ', ' ', ' ', '#', '#', '#'},
		{' ', '#', ' ', ' ', '#', ' ', ' ', '#', ' ', ' ', '#', ' ', ' ', '#', ' ', ' ', '#', ' ', ' ', ' '},
	}

	var monsterPounds int = 15

	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("You must supply a file to process.")
	}
	filename := args[0]

	tiles := processFile(filename)

	edges := findEdges(tiles)

	image := createImage(tiles, edges)

	monsters := countMonsters(image, monster, monsterPounds)
	fmt.Println("Non Monster Pounds:", monsters)
}
