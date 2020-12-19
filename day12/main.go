package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	instructions := []Instruction{}

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		instructions = append(instructions, parse(sc.Text()))
	}

	ship := &Ship{
		Direction: East,
		X:         0,
		Y:         0,
	}

	for _, i := range instructions {
		i.run(ship)
		fmt.Printf("%+v\n", i)
		fmt.Printf("%+v\n", ship)
	}

	distance := math.Abs(float64(ship.X)) + math.Abs(float64(ship.Y))
	fmt.Println(distance)
}

const (
	East    = "E"
	West    = "W"
	North   = "N"
	South   = "S"
	Forward = "F"
	Left    = "L"
	Right   = "R"
)

type Ship struct {
	Direction string
	X         int
	Y         int
}

type Instruction struct {
	Action string
	Value  int
}

func (i *Instruction) run(s *Ship) {
	switch i.Action {
	case North:
		s.Y = s.Y + i.Value
	case South:
		s.Y = s.Y - i.Value
	case East:
		s.X = s.X + i.Value
	case West:
		s.X = s.X - i.Value
	case Forward:
		switch s.Direction {
		case North:
			s.Y = s.Y + i.Value
		case South:
			s.Y = s.Y - i.Value
		case East:
			s.X = s.X + i.Value
		case West:
			s.X = s.X - i.Value
		}
	case Left:
	case Right:
		deg := deg(s.Direction)
		s.Direction = direction((deg + i.Value) % 360)
	}
}

func direction(deg int) string {
	switch deg {
	case 0:
		return East
	case 90:
		return South
	case 180:
		return West
	case 270:
		return North
	case 360:
		return East
	}
	return East
}

func deg(d string) int {
	switch d {
	case North:
		return 270
	case South:
		return 90
	case East:
		return 0
	case West:
		return 180
	}
	return 0
}

func parse(s string) Instruction {
	a := string(s[0])
	v, _ := strconv.Atoi(s[1:])
	return Instruction{
		Action: a,
		Value:  v,
	}
}
