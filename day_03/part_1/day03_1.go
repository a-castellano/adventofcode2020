// Álvaro Castellano Vela 2020/12/03
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func processFile(filename string) ([][]rune, int, int) {

	var field [][]rune

	var rows int = 0
	var columns int = 0

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		columns = len(line)
		field = append(field, line)
		rows++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return field, rows, columns

}

func countTress(field [][]rune, rows int, columns int) int {
	var trees int = 0

	var col_counter = 0

	for i := 0; i < rows; i++ {
		if field[i][col_counter%columns] == '#' {
			trees++
		}
		col_counter += 3
	}

	return trees
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("You must supply a file to process.")
	}
	filename := args[0]
	field, rows, columns := processFile(filename)

	trees := countTress(field, rows, columns)

	fmt.Println("Trees: ", trees)
}
