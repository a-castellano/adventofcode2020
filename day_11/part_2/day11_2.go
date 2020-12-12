// Álvaro Castellano Vela 2020/12/12

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
	State        SeatState
	VisibleSeats []*Seat
}

func calculateVisibleSeats(waitingArea [][]rune, width int, height int, i int, j int) []int {
	var visibleSeats []int

	//Check all directions
	var posI int
	var posJ int

	//N
	posI = i - 1
	for posI >= 0 {
		if waitingArea[posI][j] != '.' {
			visibleSeats = append(visibleSeats, (posI)*height+j)
			break
		}
		posI--
	}
	//NE
	posI = i - 1
	posJ = j + 1
	for posI >= 0 && posJ < width {
		if waitingArea[posI][posJ] != '.' {
			visibleSeats = append(visibleSeats, (posI)*height+posJ)
			break
		}
		posI--
		posJ++
	}
	//E
	posJ = j + 1
	for posJ < width {
		if waitingArea[i][posJ] != '.' {
			visibleSeats = append(visibleSeats, (i)*height+posJ)
			break
		}
		posJ++
	}
	//SE
	posI = i + 1
	posJ = j + 1
	for posI < height && posJ < width {
		if waitingArea[posI][posJ] != '.' {
			visibleSeats = append(visibleSeats, (posI)*height+posJ)
			break
		}
		posI++
		posJ++
	}
	//S
	posI = i + 1
	for posI < height {
		if waitingArea[posI][j] != '.' {
			visibleSeats = append(visibleSeats, (posI)*height+j)
			break
		}
		posI++
	}
	//SW
	posI = i + 1
	posJ = j - 1
	for posI < height && posJ >= 0 {
		if waitingArea[posI][posJ] != '.' {
			visibleSeats = append(visibleSeats, (posI)*height+posJ)
			break
		}
		posI++
		posJ--
	}
	//W
	posJ = j - 1
	for posJ >= 0 {
		if waitingArea[i][posJ] != '.' {
			visibleSeats = append(visibleSeats, (i)*height+posJ)
			break
		}
		posJ--
	}
	//NW
	posI = i - 1
	posJ = j - 1
	for posI >= 0 && posJ >= 0 {
		if waitingArea[posI][posJ] != '.' {
			visibleSeats = append(visibleSeats, (posI)*height+posJ)
			break
		}
		posI--
		posJ--
	}

	return visibleSeats
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
				visibleSeats := calculateVisibleSeats(waitingArea, width, height, i, j)
				visibleSeatsPointers := make([]*Seat, len(visibleSeats))
				for i, seatID := range visibleSeats {
					visibleSeatsPointers[i] = seatIDs[seatID]
				}
				seatIDs[i*height+j].VisibleSeats = visibleSeatsPointers
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
				var allVisibleSeatFree bool = true
				for _, visibleSeat := range seat.VisibleSeats {
					if visibleSeat.State == Occupied {
						allVisibleSeatFree = false
					}
				}
				if allVisibleSeatFree == true {
					newstates[seatIndex] = Occupied
					thisRoundHasChanges = true
					occupiedSeats++
				} else {
					newstates[seatIndex] = Free
				}
			} else { // seat.State == Occupied
				//If a seat is occupied (#) and five or more seats visible to it are also occupied, the seat becomes empty.
				var occupiedVisibleSeats int = 0
				for _, visibleSeat := range seat.VisibleSeats {
					if visibleSeat.State == Occupied {
						occupiedVisibleSeats++
					}
				}
				if occupiedVisibleSeats >= 5 {
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
