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
		switch s.Direction {
		case East:
			switch i.Value {
			case 90:
				s.Direction = North
			case 180:
				s.Direction = West
			case 270:
				s.Direction = South
			}
		case West:
			switch i.Value {
			case 90:
				s.Direction = South
			case 180:
				s.Direction = East
			case 270:
				s.Direction = North
			}
		case North:
			switch i.Value {
			case 90:
				s.Direction = West
			case 180:
				s.Direction = South
			case 270:
				s.Direction = East
			}
		case South:
			switch i.Value {
			case 90:
				s.Direction = East
			case 180:
				s.Direction = North
			case 270:
				s.Direction = West
			}
		}
	case Right:
		switch s.Direction {
		case East:
			switch i.Value {
			case 90:
				s.Direction = South
			case 180:
				s.Direction = West
			case 270:
				s.Direction = North
			}
		case West:
			switch i.Value {
			case 90:
				s.Direction = North
			case 180:
				s.Direction = East
			case 270:
				s.Direction = South
			}
		case North:
			switch i.Value {
			case 90:
				s.Direction = East
			case 180:
				s.Direction = South
			case 270:
				s.Direction = West
			}
		case South:
			switch i.Value {
			case 90:
				s.Direction = West
			case 180:
				s.Direction = North
			case 270:
				s.Direction = East
			}
		}
	}
}

func parse(s string) Instruction {
	a := string(s[0])
	v, _ := strconv.Atoi(s[1:])
	return Instruction{
		Action: a,
		Value:  v,
	}
}
