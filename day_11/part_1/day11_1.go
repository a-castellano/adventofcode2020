// Álvaro Castellano Vela 2020/12/11
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type SeatState uint8

const (
	Free SeatState = 1 << iota
	Occupied
)

type Seat struct {
	State         SeatState
	AdjacentSeats []*Seat
}

func calculateAdjacentSeats(waitingArea [][]rune, width int, height int, i int, j int) []int {

	var ajacentSeats []int

	if i == 0 {
		if j == 0 {
			if waitingArea[i][j+1] == 'L' {
				ajacentSeats = append(ajacentSeats, i*height+j+1)
			}
			if waitingArea[i+1][j] == 'L' {
				ajacentSeats = append(ajacentSeats, (i+1)*height+j)
			}
			if waitingArea[i+1][j+1] == 'L' {
				ajacentSeats = append(ajacentSeats, (i+1)*height+j)
			}
		} else if j != width-1 {
			if waitingArea[i][j-1] == 'L' {
				ajacentSeats = append(ajacentSeats, i*height+j-1)
			}
			if waitingArea[i][j+1] == 'L' {
				ajacentSeats = append(ajacentSeats, i*height+j+1)
			}
			if waitingArea[i+1][j-1] == 'L' {
				ajacentSeats = append(ajacentSeats, (i+1)*height+j-1)
			}
			if waitingArea[i+1][j] == 'L' {
				ajacentSeats = append(ajacentSeats, (i+1)*height+j)
			}
			if waitingArea[i+1][j+1] == 'L' {
				ajacentSeats = append(ajacentSeats, (i+1)*height+j+1)
			}

		} else { // j == width-1
			if waitingArea[i][j-1] == 'L' {
				ajacentSeats = append(ajacentSeats, i*height+j-1)
			}
			if waitingArea[i+1][j-1] == 'L' {
				ajacentSeats = append(ajacentSeats, (i+1)*height+j-1)
			}
			if waitingArea[i+1][j] == 'L' {
				ajacentSeats = append(ajacentSeats, (i+1)*height+j)
			}
		}
	} else if i != height-1 {

		if j == 0 {
			if waitingArea[i-1][j] == 'L' {
				ajacentSeats = append(ajacentSeats, (i-1)*height+j)
			}
			if waitingArea[i-1][j+1] == 'L' {
				ajacentSeats = append(ajacentSeats, (i-1)*height+j+1)
			}
			if waitingArea[i][j+1] == 'L' {
				ajacentSeats = append(ajacentSeats, i*height+j+1)
			}
			if waitingArea[i+1][j+1] == 'L' {
				ajacentSeats = append(ajacentSeats, (i+1)*height+j+1)
			}
			if waitingArea[i+1][j] == 'L' {
				ajacentSeats = append(ajacentSeats, (i+1)*height+j)
			}
		} else if j != width-1 {
			if waitingArea[i-1][j-1] == 'L' {
				ajacentSeats = append(ajacentSeats, (i-1)*height+j-1)
			}
			if waitingArea[i-1][j] == 'L' {
				ajacentSeats = append(ajacentSeats, (i-1)*height+j)
			}
			if waitingArea[i-1][j+1] == 'L' {
				ajacentSeats = append(ajacentSeats, (i-1)*height+j+1)
			}
			if waitingArea[i][j-1] == 'L' {
				ajacentSeats = append(ajacentSeats, i*height+j-1)
			}
			if waitingArea[i][j+1] == 'L' {
				ajacentSeats = append(ajacentSeats, i*height+j+1)
			}
			if waitingArea[i+1][j-1] == 'L' {
				ajacentSeats = append(ajacentSeats, (i+1)*height+j-1)
			}
			if waitingArea[i+1][j] == 'L' {
				ajacentSeats = append(ajacentSeats, (i+1)*height+j)
			}
			if waitingArea[i+1][j+1] == 'L' {
				ajacentSeats = append(ajacentSeats, (i+1)*height+j+1)
			}
		} else { // j == width-1
			if waitingArea[i-1][j-1] == 'L' {
				ajacentSeats = append(ajacentSeats, (i-1)*height+j-1)
			}
			if waitingArea[i-1][j] == 'L' {
				ajacentSeats = append(ajacentSeats, (i-1)*height+j)
			}
			if waitingArea[i][j-1] == 'L' {
				ajacentSeats = append(ajacentSeats, i*height+j-1)
			}
			if waitingArea[i+1][j-1] == 'L' {
				ajacentSeats = append(ajacentSeats, (i+1)*height+j-1)
			}
			if waitingArea[i+1][j] == 'L' {
				ajacentSeats = append(ajacentSeats, (i+1)*height+j)
			}
		}
	} else { // i == height-1
		if j == 0 {
			if waitingArea[i-1][j] == 'L' {
				ajacentSeats = append(ajacentSeats, (i-1)*height+j)
			}
			if waitingArea[i-1][j+1] == 'L' {
				ajacentSeats = append(ajacentSeats, (i-1)*height+j+1)
			}
			if waitingArea[i][j+1] == 'L' {
				ajacentSeats = append(ajacentSeats, (i)*height+j+1)
			}
		} else if j != width-1 {
			if waitingArea[i][j-1] == 'L' {
				ajacentSeats = append(ajacentSeats, i*height+j-1)
			}
			if waitingArea[i-1][j-1] == 'L' {
				ajacentSeats = append(ajacentSeats, (i-1)*height+j-1)
			}
			if waitingArea[i-1][j] == 'L' {
				ajacentSeats = append(ajacentSeats, (i-1)*height+j)
			}
			if waitingArea[i-1][j+1] == 'L' {
				ajacentSeats = append(ajacentSeats, (i-1)*height+j+1)
			}
			if waitingArea[i][j+1] == 'L' {
				ajacentSeats = append(ajacentSeats, i*height+j+1)
			}
		} else { // j == width-1
			if waitingArea[i-1][j] == 'L' {
				ajacentSeats = append(ajacentSeats, (i-1)*height+j)
			}
			if waitingArea[i-1][j-1] == 'L' {
				ajacentSeats = append(ajacentSeats, (i-1)*height+j-1)
			}
			if waitingArea[i-1][j] == 'L' {
				ajacentSeats = append(ajacentSeats, (i-1)*height+j)
			}
		}
	}

	return ajacentSeats
}

func processFile(filename string) []*Seat {

	var seats []*Seat
	var waitingArea [][]rune
	seatIDs := make(map[int]*Seat)

	var width int = 0
	var height int = 0

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		row := []rune(scanner.Text())
		width = len(row)
		waitingArea = append(waitingArea, row)
		height++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// create seats
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if waitingArea[i][j] == 'L' {
				seat := Seat{Free, make([]*Seat, 0)}
				seatIDs[i*height+j] = &seat
				seats = append(seats, &seat)
			}
		}
	}
	// calculate adjacent
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if waitingArea[i][j] == 'L' {
				adjacentSeats := calculateAdjacentSeats(waitingArea, width, height, i, j)
				adjacentSeatsPointers := make([]*Seat, len(adjacentSeats))
				for i, seatID := range adjacentSeats {
					adjacentSeatsPointers[i] = seatIDs[seatID]
				}
				seatIDs[i*height+j].AdjacentSeats = adjacentSeatsPointers
			}
		}
	}

	return seats
}

func chaosStabilizes(seats []*Seat) int {

	var occupiedSeatsAferChaosStabilizes int = 0
	var totalSeats int = len(seats)

	for {

		var occupiedSeats int = 0
		var thisRoundHasChanges bool = false
		newstates := make([]SeatState, totalSeats)

		for seatIndex, seat := range seats {
			//If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
			if seat.State == Free {
				var allAdjacentFree bool = true
				for _, adjacentSeat := range seat.AdjacentSeats {
					if adjacentSeat.State == Occupied {
						allAdjacentFree = false
					}
				}
				if allAdjacentFree == true {
					newstates[seatIndex] = Occupied
					thisRoundHasChanges = true
					occupiedSeats++
				} else {
					newstates[seatIndex] = Free
				}
			} else { // seat.State == Occupied
				//If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
				var occupiedAdjacentSeats int = 0
				for _, adjacentSeat := range seat.AdjacentSeats {
					if adjacentSeat.State == Occupied {
						occupiedAdjacentSeats++
					}
				}
				if occupiedAdjacentSeats >= 4 {
					newstates[seatIndex] = Free
					thisRoundHasChanges = true
				} else {
					newstates[seatIndex] = Occupied
					occupiedSeats++
				}
			}
		}

		if thisRoundHasChanges {
			//Apply changes
			for seatIndex, seat := range seats {
				seat.State = newstates[seatIndex]
			}
		} else {
			occupiedSeatsAferChaosStabilizes = occupiedSeats
			break
		}

	}
	return occupiedSeatsAferChaosStabilizes
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("You must supply a file to process.")
	}
	filename := args[0]

	seats := processFile(filename)

	occupiedSeats := chaosStabilizes(seats)
	fmt.Println("Occupied seats after chaos stabilizes:", occupiedSeats)
}
