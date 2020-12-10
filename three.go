package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var treemap [][]string

	sc := bufio.NewScanner(os.Stdin)
	x := 0
	for sc.Scan() {
		row := sc.Text()
		treemap = append(treemap, []string{})

		// repeat to expand map rightward...
		for i := 0; i < 10000; i++ {
			for _, c := range row {
				treemap[x] = append(treemap[x], string(c))
			}
		}
		x = x + 1
	}

	a := sled(1, 1, treemap)
	b := sled(3, 1, treemap)
	c := sled(5, 1, treemap)
	d := sled(7, 1, treemap)
	e := sled(1, 2, treemap)

	fmt.Println(a * b * c * d * e)
}

func sled(x int, y int, treemap [][]string) int {
	trees := 0
	mx := 0
	my := 0
	for {
		if len(treemap) > my {
			if len(treemap[my]) > mx {
				if string(treemap[my][mx]) == "#" {
					trees = trees + 1
				}
				mx = mx + x
				my = my + y
			} else {
				return trees
			}
		} else {
			return trees
		}
	}

	return 0
}
