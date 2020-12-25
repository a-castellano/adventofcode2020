// Álvaro Castellano Vela 2020/12/25
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func findLoop(publicKey int) int {

	var loopSize int = 0
	var subject int = 7
	var result int = 1

	for result != publicKey {
		result = result * subject
		result = result % 20201227
		loopSize++
	}

	return loopSize
}

func applyLoop(subject int, loopSize int) int {

	var result int = 1

	for loopSize > 0 {
		result = result * subject
		result = result % 20201227
		loopSize--
	}

	return result
}

func main() {
	//	var result int
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatal("You must supply card's public key and door's public key")
	}
	cardPublicKey, _ := strconv.Atoi(args[0])
	doorPublicKey, _ := strconv.Atoi(args[1])

	cardLoop := findLoop(cardPublicKey)
	doorLoop := findLoop(doorPublicKey)

	cardEncryption := applyLoop(cardPublicKey, doorLoop)
	doorEncryption := applyLoop(doorPublicKey, cardLoop)

	if cardEncryption == doorEncryption {
		fmt.Println("Encryption", cardEncryption)
	} else {
		log.Fatal("Failed to obtain Encryption")
	}
}
