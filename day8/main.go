package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	program := []instruction{}

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		program = append(program, parse(sc.Text()))
	}

	// fmt.Println(run(program))
	findPart2Solution(program)
}

func run(program []instruction) (int, error) {
	accumulator := 0

	history := make(map[int]bool)
	i := 0

	for {
		if history[i] {
			return accumulator, errors.New("Infinite loop detected")
		} else if i >= len(program) {
			return accumulator, nil
		}
		history[i] = true

		ins := program[i]

		switch ins.op {
		case "acc":
			accumulator = accumulator + ins.arg
			i = i + 1
		case "jmp":
			i = i + ins.arg
		case "nop":
			i = i + 1
		}
	}

}

func makePrograms(program []instruction) [][]instruction {
	result := [][]instruction{}
	result = append(result, program)

	for i, ins := range program {
		if ins.op == "jmp" {
			newProgram := make([]instruction, len(program))
			copy(newProgram, program)
			newProgram[i] = instruction{
				op:  "nop",
				arg: ins.arg,
			}
			result = append(result, newProgram)
		} else if ins.op == "nop" {
			newProgram := make([]instruction, len(program))
			copy(newProgram, program)
			newProgram[i] = instruction{
				op:  "jmp",
				arg: ins.arg,
			}
			result = append(result, newProgram)
		}
	}

	return result
}

func findPart2Solution(program []instruction) {
	programs := makePrograms(program)

	for _, p := range programs {
		acc, err := run(p)
		if err == nil {
			fmt.Println(acc)
			break
		}
	}
}

type instruction struct {
	op  string // "acc" | "jmp" | "nop"
	arg int
}

func parse(input string) instruction {
	res := strings.Split(input, " ")
	op := res[0]
	arg, _ := strconv.Atoi(res[1])

	return instruction{
		op:  op,
		arg: arg,
	}
}
