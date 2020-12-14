// Álvaro Castellano Vela 2020/12/14
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func decToBin(dec int) []uint8 {
	bin := make([]uint8, 36)
	var counter int = 35

	for dec != 0 {
		bin[counter] = uint8(dec % 2)
		dec = dec / 2
		counter--
	}

	return bin
}

func binToDec(bin []uint8) int {
	var dec int
	var power int = 1

	for pos := 35; pos >= 0; pos-- {
		dec += int(bin[pos]) * power
		power *= 2
	}
	return dec
}

func calculateValue(value []uint8, mask []uint8) []uint8 {
	newValue := make([]uint8, 36)

	for pos := 0; pos < 36; pos++ {
		if mask[pos] == 2 { // X
			newValue[pos] = value[pos]
		} else {
			newValue[pos] = mask[pos]
		}
	}
	return newValue
}

func calculateNewAddressBase(address []uint8, bitmask []uint8) []uint8 {
	newAddress := make([]uint8, 36)

	for pos, _ := range address {
		if bitmask[pos] == 0 {
			newAddress[pos] = address[pos]
		} else {
			newAddress[pos] = bitmask[pos]
		}
	}

	return newAddress
}

func calculateMemoryPermutations(memoryPermutations map[int]bool, address []uint8, position int) {
	for position > 0 && address[position] != 2 {
		position--
	}
	if address[position] != 2 { //final
		memoryPermutations[binToDec(address)] = true
	} else { // split

		firstOption := make([]uint8, 36)
		secondOption := make([]uint8, 36)

		copy(firstOption, address)
		copy(secondOption, address)
		firstOption[position] = 0
		secondOption[position] = 1

		calculateMemoryPermutations(memoryPermutations, firstOption, position)
		calculateMemoryPermutations(memoryPermutations, secondOption, position)
	}
}

func processCode(filename string) map[int][]uint8 {

	memory := make(map[int][]uint8)

	currentBitmask := make([]uint8, 36)

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	maskRule := regexp.MustCompile(`^mask = ([X01]{36})$`)
	writeRule := regexp.MustCompile(`^mem\[([0-9]+)\] = ([0-9]+)$`)

	for scanner.Scan() {
		matchMask := maskRule.FindAllStringSubmatch(scanner.Text(), -1)
		matchWrite := writeRule.FindAllStringSubmatch(scanner.Text(), -1)
		if len(matchMask) != 0 { // Found bitmask
			for pos, bitValue := range matchMask[0][1] {
				if bitValue == 'X' {
					currentBitmask[pos] = 2
				} else if bitValue == '0' {
					currentBitmask[pos] = 0
				} else {
					currentBitmask[pos] = 1
				}
			}
		} else { // Found Write
			memoryPosition, _ := strconv.Atoi(matchWrite[0][1])
			if _, ok := memory[memoryPosition]; !ok {
				newMemoryPostion := make([]uint8, 36)
				memory[memoryPosition] = newMemoryPostion
			}
			value, _ := strconv.Atoi(matchWrite[0][2])

			newAddressBase := calculateNewAddressBase(decToBin(memoryPosition), currentBitmask)
			memoryPermutations := make(map[int]bool)

			calculateMemoryPermutations(memoryPermutations, newAddressBase, 35)
			newValue := decToBin(value)
			for permutatedMemoryPosition, _ := range memoryPermutations {
				memory[permutatedMemoryPosition] = newValue
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return memory
}

func sumMemory(memory map[int][]uint8) int {
	var result = 0
	for _, value := range memory {
		result += binToDec(value)
	}

	return result
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("You must supply a file to process.")
	}
	filename := args[0]

	memory := processCode(filename)
	result := sumMemory(memory)
	fmt.Println("Result:", result)
}
