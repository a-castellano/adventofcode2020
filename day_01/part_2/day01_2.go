// Álvaro Castellano Vela 2020/12/01
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func processFile(filename string) []int {

	var entries []int

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		entry, _ := strconv.Atoi(scanner.Text())
		entries = append(entries, entry)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return entries
}

func calculteResult(entries []int) int {

	var result int
	var found bool = false

	for i, entry1 := range entries {
		for j, entry2 := range entries {
			for k, entry3 := range entries {
				if i != j && j != k && i != k {
					if entry1+entry2+entry3 == 2020 {
						result = entry1 * entry2 * entry3
						found = true
						break
					}
				}
			}
			if found {
				break
			}

		}
		if found {
			break
		}
	}

	return result

}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("You must supply a file to process.")
	}
	filename := args[0]
	entries := processFile(filename)
	result := calculteResult(entries)
	fmt.Printf("Result: %d\n", result)
}
