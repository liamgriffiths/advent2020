package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"

	// "reflect"
	"strings"
)

const (
	Empty    = "L"
	Occupied = "#"
	Floor    = "."
)

func main() {
	n := 0
	seats := [][]string{}

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := sc.Text()
		seats = append(seats, []string{})
		chairs := strings.Split(line, "")
		for _, c := range chairs {
			seats[n] = append(seats[n], c)
		}
		n = n + 1
	}

	var prev [][]string
	var next [][]string

	prev = seats
	for {
		next = iterate(prev)

		if reflect.DeepEqual(prev, next) {
			fmt.Println("finalCount:", finalCount(next))
			break
		}

		prev = next
	}
}

func iterate(seats [][]string) [][]string {
	result := [][]string{}

	// all the directions for adj seats
	dirs := [][]int{
		{-1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
	}

	for x, _ := range seats {
		result = append(result, []string{})
		for y, _ := range seats[x] {
			adj := []string{}
			for _, d := range dirs {
				ax, ay := x+d[0], y+d[1]
				if ax < 0 {
					continue
				}
				if ay < 0 {
					continue
				}
				if ax >= len(seats) {
					continue
				}
				if ay >= len(seats[x]) {
					continue
				}

				adj = append(adj, seats[ax][ay])
			}

			next := rule(seats[x][y], adj)
			result[x] = append(result[x], next)
		}
	}
	return result
}

func rule(seat string, adjSeats []string) string {
	switch seat {
	case Empty:
		if occupiedCount(adjSeats) == 0 {
			return Occupied
		}
	case Occupied:
		if occupiedCount(adjSeats) >= 4 {
			return Empty
		}
	case Floor:
		// do nothing
	}
	return seat
}

func finalCount(seats [][]string) int {
	res := []string{}
	for x, _ := range seats {
		for y, _ := range seats[x] {
			res = append(res, seats[x][y])
		}
	}
	return occupiedCount(res)
}

func occupiedCount(seats []string) int {
	c := 0
	for _, v := range seats {
		if v == Occupied {
			c = c + 1
		}
	}
	return c
}

func pretty(seats [][]string) {
	for x, _ := range seats {
		for y, _ := range seats[x] {
			fmt.Print(seats[x][y])
		}
		fmt.Println("")
	}
	fmt.Println("")
}
