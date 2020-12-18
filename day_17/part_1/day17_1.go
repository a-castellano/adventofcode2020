// Álvaro Castellano Vela 2020/12/17
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	//"strings"
)

type Cube struct {
	Active    bool
	X         int
	Y         int
	Z         int
	Neighbors []*Cube
}

func (cube *Cube) SetNeighbors(cubes map[string]*Cube, neighbors map[string]map[string]*Cube) {
	var cubeStringCoordinates string = StringCoordinates(cube.X, cube.Y, cube.Z)
	var stringNeighborCoordinates string
	cubeNeighbors := make(map[string]*Cube)
	count := 0
	for z := -1; z <= 1; z++ {
		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {
				stringNeighborCoordinates = StringCoordinates(cube.X+x, cube.Y+y, cube.Z+z)
				if stringNeighborCoordinates != cubeStringCoordinates {
					var neighbor Cube
					if _, ok := cubes[stringNeighborCoordinates]; !ok {
						neighbor.X = x
						neighbor.Y = y
						neighbor.Z = z
						neighbor.Active = false
					}
					cubeNeighbors[stringNeighborCoordinates] = cubes[stringNeighborCoordinates]
					count++
				}
			}
		}
	}
	neighbors[cubeStringCoordinates] = cubeNeighbors
	cube.HasNeighbors = true
}

func StringCoordinates(X, Y, Z int) string {
	return fmt.Sprintf("%d,%d,%d", X, Y, Z)
}

func processFile(filename string) (map[string]*Cube, map[string]map[string]*Cube) {

	cubes := make(map[string]*Cube)
	neighbors := make(map[string]map[string]*Cube)

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
			cube := Cube{false, column, row, 0, false}
			if character == '#' { // Active
				cube.Active = true
			}
			cubes[StringCoordinates(cube.X, cube.Y, cube.Z)] = &cube
			column++
		}
		row++
	}

	fmt.Println(len(cubes))

	for _, cube := range cubes {
		cube.SetNeighbors(cubes, neighbors)
	}

	return cubes, neighbors
}

func runCicles(cubes map[string]*Cube, neighbors map[string]map[string]*Cube, cicles int) {
	newStates := make(map[string]bool)
	for _, cube := range cubes {
		var activeCubes int = 0
		var stateChanged bool = false
		var cubeStringCoordinates = StringCoordinates(cube.X, cube.Y, cube.Z)
		newStates[cubeStringCoordinates] = cube.Active
		for _, neighbor := range neighbors[cubeStringCoordinates] {
			if neighbor.Active == true {
				activeCubes++
			}

		}

		if cube.Active {
			fmt.Println("Cube ", cubeStringCoordinates, "has ", activeCubes, "active neighbors")
			if activeCubes != 2 && activeCubes != 3 {
				newStates[cubeStringCoordinates] = false
				stateChanged = true
			} else {
				newStates[cubeStringCoordinates] = true
			}
		} else {
			if cube.Active == false && activeCubes == 3 {
				newStates[cubeStringCoordinates] = true
				stateChanged = true
			} else {
				newStates[cubeStringCoordinates] = false
			}
		}
		if stateChanged && cube.HasNeighbors == false {
			fmt.Println("Changing Neighbors of", cube)
			cube.SetNeighbors(cubes, neighbors)
			fmt.Println("after Neighbors of", cube)
		}
	}
	for cubeStringCoordinates, _ := range cubes {
		cubes[cubeStringCoordinates].Active = newStates[cubeStringCoordinates]
	}
}

func countActive(cubes map[string]*Cube) int {
	var activeCubes int = 0
	for _, cube := range cubes {
		fmt.Println(cube)
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

	cubes, neighbors := processFile(filename)
	fmt.Println("Cubes", cubes)
	fmt.Println("Neighbors", neighbors)
	fmt.Println(len(cubes))
	fmt.Println("activeCubes:", countActive(cubes))
	fmt.Println("__________")
	runCicles(cubes, neighbors, 1)
	fmt.Println("Cubes", cubes)
	fmt.Println("Neighbors", neighbors)
	fmt.Println("__________")
	fmt.Println(len(cubes))
	fmt.Println("activeCubes:", countActive(cubes))
}
