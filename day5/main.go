package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	input := []string{}

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		input = append(input, sc.Text())
	}

	bps := []*boardingPass{}
	for _, in := range input {
		bps = append(bps, parse(in))
	}

	// fmt.Println(highestSeatId(bps))
	fmt.Println(findSeat(bps))
}

func highestSeatId(bps []*boardingPass) int {
	highest := 0
	for _, bp := range bps {
		if bp.id > highest {
			highest = bp.id
		}
	}
	return highest
}

func allSeats() map[int]bool {
	seats := make(map[int]bool)

	for r := 0; r < 128; r++ {
		for c := 0; c < 8; c++ {
			id := r*8 + c
			seats[id] = false
		}
	}

	return seats
}

func findEmptySeats(bps []*boardingPass) ([]int, map[int]bool) {
	emptySeatIds := []int{}
	seats := allSeats()

	for _, bp := range bps {
		seats[bp.id] = true
	}

	for k, v := range seats {
		if !v {
			emptySeatIds = append(emptySeatIds, k)
		}
	}

	return emptySeatIds, seats
}

func findSeat(bps []*boardingPass) int {
	emptyIds, seats := findEmptySeats(bps)
	sort.Ints(emptyIds)

	for i, v := range emptyIds {
		if i > 0 {
			b := seats[v-1]
			a := seats[v+1]

			if b && a {
				return v
			}
		}
	}

	return 0
}

type boardingPass struct {
	code string
	row  int // 0 - 127
	col  int // 0 - 7
	id   int // row * 8 + col
}

func parse(code string) *boardingPass {
	rows := []int{}
	for i := 0; i < 128; i++ {
		rows = append(rows, i)
	}

	cols := []int{}
	for i := 0; i < 8; i++ {
		cols = append(cols, i)
	}

	for _, c := range code {
		switch string(c) {
		case "F":
			// lower half (rows)
			rows = rows[:len(rows)/2]
		case "B":
			// upper half (rows)
			rows = rows[len(rows)/2:]
		case "L":
			// lower half (cols)
			cols = cols[:len(cols)/2]
		case "R":
			// upper half (cols)
			cols = cols[len(cols)/2:]
		}
	}

	return &boardingPass{
		code: code,
		row:  rows[0],
		col:  cols[0],
		id:   rows[0]*8 + cols[0],
	}
}
