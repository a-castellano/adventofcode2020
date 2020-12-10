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

func countDifferences(adapters []int) (int, int) {
	var difference1 int = 0
	var difference3 int = 0

	for i := 0; i < len(adapters)-1; i++ {
		difference := adapters[i+1] - adapters[i]
		if difference == 1 {
			difference1++
		} else if difference == 3 {
			difference3++
		}
	}
	// our our device's built-in adapter
	difference3++

	return difference1, difference3

}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("You must supply a file to process.")
	}
	filename := args[0]

	adapters := processFile(filename)
	difference1, difference3 := countDifferences(adapters)
	fmt.Println("Result:", difference1*difference3)
}
