// Álvaro Castellano Vela 2020/12/10
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func processFile(filename string) []int {

	var adapters []int

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		adapter, _ := strconv.Atoi(scanner.Text())
		adapters = append(adapters, adapter)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// add charging outlet
	adapters = append(adapters, 0)
	sort.Ints(adapters)

	return adapters
}

func get3DifferencesArray(adapters []int) []int {

	var difference3Array []int

	difference3Array = append(difference3Array, 0)
	for i := 0; i < len(adapters)-1; i++ {
		difference := adapters[i+1] - adapters[i]
		if difference == 3 {
			if difference3Array[len(difference3Array)-1] != i {
				difference3Array = append(difference3Array, i)
			}
			difference3Array = append(difference3Array, i+1)
		}
	}
	//Last adapters must be added always
	difference3Array = append(difference3Array, len(adapters)-1)
	return difference3Array

}

func Pow(n int) int {
	var result = 2
	for i := n; i > 1; i-- {
		result *= 2
	}
	return result
}

func calculatePermutations(difference3Array []int, adapters []int) int {
	var permutations int = 1
	for index := 0; index < len(difference3Array); index++ {
		permutableNumbers := difference3Array[(index+1)%len(difference3Array)] - difference3Array[index] - 1
		if permutableNumbers > 0 {
			if permutableNumbers == 3 {
				permutations *= 7
			} else if permutableNumbers == 2 {
				permutations *= 4
			} else if permutableNumbers == 1 {
				permutations *= 2
			}
		}

	}

	return permutations
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("You must supply a file to process.")
	}
	filename := args[0]

	adapters := processFile(filename)
	difference3Array := get3DifferencesArray(adapters)
	permutations := calculatePermutations(difference3Array, adapters)
	fmt.Println("Permutations:", permutations)
}
