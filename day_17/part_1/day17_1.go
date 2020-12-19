// Álvaro Castellano Vela 2020/12/19
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Cube struct {
	Active             bool
	X                  int
	Y                  int
	Z                  int
	Neighbors          []*Cube
	AllocatedNeighbors int
}

func StringCoordinates(x, y, z int) string {
	return fmt.Sprintf("%d,%d,%d", x, y, z)
}

func (cube *Cube) SetPartialNeighbors(cubes map[string]*Cube) {
	var stringNeighborCoordinates string
	var neighborCounter int = 0
	for z := -1; z <= 1; z++ {
		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {
				if !(x == 0 && y == 0 && z == 0) {
					stringNeighborCoordinates = StringCoordinates(cube.X+x, cube.Y+y, cube.Z+z)
					if _, ok := cubes[stringNeighborCoordinates]; ok {
						cube.Neighbors[neighborCounter] = cubes[stringNeighborCoordinates]
						neighborCounter++
					}
				}
			}
		}
	}
	cube.AllocatedNeighbors = neighborCounter
}

func (cube *Cube) SetNeighbors(cubes map[string]*Cube, neighbors map[string]*Cube) {
	var stringNeighborCoordinates string
	var neighborCounter int = 0
	cube.AllocatedNeighbors = 0
	for z := -1; z <= 1; z++ {
		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {
				if !(x == 0 && y == 0 && z == 0) {
					stringNeighborCoordinates = StringCoordinates(cube.X+x, cube.Y+y, cube.Z+z)
					if _, incubes := cubes[stringNeighborCoordinates]; !incubes {
						if _, inneighbors := neighbors[stringNeighborCoordinates]; !inneighbors {
							neighbor := Cube{false, cube.X + x, cube.Y + y, cube.Z + z, make([]*Cube, 26), 0}
							cubes[stringNeighborCoordinates] = &neighbor
							neighbors[stringNeighborCoordinates] = &neighbor
						}
					}
					cube.Neighbors[neighborCounter] = cubes[stringNeighborCoordinates]
					neighborCounter++
				}
			}
		}
	}
	cube.AllocatedNeighbors = neighborCounter
}

func getCurrentCubes(cubes map[string]*Cube) []string {
	cubeList := make([]string, len(cubes))
	var counter int = 0
	for stringCoordinates, _ := range cubes {
		cubeList[counter] = stringCoordinates
		counter++
	}
	return cubeList
}

func processFile(filename string) map[string]*Cube {

	cubes := make(map[string]*Cube)
	neighbors := make(map[string]*Cube)

	var row int = 0

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// read map
	for scanner.Scan() {
		var column int = 0
		for _, character := range scanner.Text() {
			cube := Cube{false, column, row, 0, make([]*Cube, 26), 0}
			if character == '#' { // Active
				cube.Active = true
			}
			cubes[StringCoordinates(cube.X, cube.Y, cube.Z)] = &cube
			column++
		}
		row++
	}

	counter := 0
	currentCubes := getCurrentCubes(cubes)
	for _, stringCoordinates := range currentCubes {
		cubes[stringCoordinates].SetNeighbors(cubes, neighbors)
		counter++
	}
	currentNeighbors := getCurrentCubes(neighbors)
	for _, stringCoordinates := range currentNeighbors {
		neighbors[stringCoordinates].SetNeighbors(cubes, neighbors)
	}

	return cubes
}

func runCicles(cubes map[string]*Cube, cicles int) {
	for cicle := 0; cicle < cicles; cicle++ {
		newStates := make(map[string]bool)
		neighbors := make(map[string]*Cube)
		for _, cube := range cubes {
			var activeCubes int = 0
			var cubeStringCoordinates = StringCoordinates(cube.X, cube.Y, cube.Z)
			newStates[cubeStringCoordinates] = cube.Active
			for neighbor := 0; neighbor < cube.AllocatedNeighbors; neighbor++ {
				if cube.Neighbors[neighbor].Active == true {
					activeCubes++
				}
			}

			if cube.Active {
				//	fmt.Println("Cube ", cubeStringCoordinates, "has ", activeCubes, "active neighbors")
				if activeCubes != 2 && activeCubes != 3 {
					newStates[cubeStringCoordinates] = false
					//	stateChanged = true
				} else {
					newStates[cubeStringCoordinates] = true
				}
			} else {
				if cube.Active == false && activeCubes == 3 {
					newStates[cubeStringCoordinates] = true
					//		stateChanged = true
				} else {
					newStates[cubeStringCoordinates] = false
				}
			}
			if cube.AllocatedNeighbors != 26 {
				cube.SetNeighbors(cubes, neighbors)
			}
		}
		for cubeStringCoordinates, cube := range cubes {
			cube.Active = newStates[cubeStringCoordinates]
		}
		currentNeighbors := getCurrentCubes(neighbors)
		for _, stringCoordinates := range currentNeighbors {
			neighbors[stringCoordinates].SetNeighbors(cubes, neighbors)
		}

	}
}

func countActive(cubes map[string]*Cube) int {
	var activeCubes int = 0
	for _, cube := range cubes {
		if cube.Active {
			activeCubes++
		}
	}
	return activeCubes
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("You must supply a file to process.")
	}
	filename := args[0]

	cubes := processFile(filename)
	runCicles(cubes, 6)
	fmt.Println("Active Cubes:", countActive(cubes))
}
