package main

import (
	"fmt"
	"log"
	"strings"
	// Used for user auth
)

const taskURL = "https://adventofcode.com/2020/day/11/input"

var wrWidth int

func main() {
	data, err := getData(taskURL)
	if err != nil {
		log.Printf("Error fetching data: %v", err)
	}

	wr := new(waitingRoom)
	wr.rows = strings.Split(data, "\n")
	wrWidth = len(wr.rows[0])

	// Strip end line for correct padding
	wr.rows = wr.rows[:len(wr.rows)-1]
	wr.padDataWithFloor()

	newWr := new(waitingRoom)

	for {
		if wr != newWr {
			// Cycle through each seat and apply rules
			for i := 1; i < len(wr.rows)-1; i++ {
				newString := make([]string, wrWidth)
				for j := 1; j < len(wr.rows[i])-1; j++ {
					mySeat := new(seat)
					mySeat.x = i
					mySeat.y = j
					mySeat.content = mySeat.flipSeat(*wr)
					newString = append(newString, mySeat.content)
				}
				newWr.rows = append(newWr.rows, strings.Join(newString, ""))
			}
		}
		break
	}

	// for i := range newWr.rows {
	// 	fmt.Printf("%v", newWr.rows[i])
	// }
}

type waitingRoom struct {
	rows []string
}

type seat struct {
	content string
	position
}

type position struct {
	x int
	y int
}

func (s *seat) flipSeat(wr waitingRoom) string {
	// If a seat is empty (L) and there are no occupied seats adjacent to it,
	// the seat becomes occupied.
	// If a seat is occupied (#) and four or more seats adjacent to it are also occupied,
	// the seat becomes empty.
	// Otherwise, the seat's state does not change.

	if string(wr.rows[s.x][s.y]) == "." {
		return "."
	}
	fmt.Printf("Seat x: %v y: %v\n", s.x, s.y)
	n := make([]string, 9)

	// Top
	n[0] = wr.rows[(s.x)-1]

	n[1] = wr.rows[(s.x)-1]
	n[2] = wr.rows[(s.x+1)-1]

	// Middle
	n[3] = wr.rows[(s.x - 1)]
	n[4] = "X" // We will dicard this
	n[5] = wr.rows[(s.x + 1)]

	// Bottom
	n[6] = wr.rows[(s.x-1)+1]
	n[7] = wr.rows[(s.x)+1]
	n[8] = wr.rows[(s.x+1)+1]

	var seatCount string
	for i := 0; i < 9; i++ {
		seatCount = seatCount + n[i]
	}

	seatHash := strings.Count(seatCount, "#")

	if s.content == "L" {
		if seatHash == 0 {
			fmt.Println("Flipped!")
			return "#"
		}
	}
	if s.content == "#" {
		if seatHash >= 4 {
			fmt.Println("Flipped!")
			return "L"
		}
	}
	return "."
}

// padDataWithFloor will allow easier parsing for the rules
// by adding a border of floor (.) characters.
func (d *waitingRoom) padDataWithFloor() {
	paddedData := waitingRoom{}

	// Pad top
	topAndBottom := strings.Repeat(".", (wrWidth+2)) + "\n"
	paddedData.rows = append(paddedData.rows, topAndBottom)

	// Pad sides
	for i := range d.rows {
		paddedData.rows = append(paddedData.rows, "."+d.rows[i]+".\n")
	}

	// Pad bottom
	paddedData.rows = append(paddedData.rows, topAndBottom)

	*d = paddedData
}
