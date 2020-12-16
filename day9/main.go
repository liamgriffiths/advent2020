package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	input := []int{}
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		i, _ := strconv.Atoi(sc.Text())
		input = append(input, i)
	}

	// sample
	// fmt.Println(xmas(input, 5))

	// part 1
	num := xmas(input, 25)
	fmt.Println(num)
	// part 2
	w := weakness(input, num)
	fmt.Println(w)
}

// find the first number in the list (after the preamble) which is not the sum
// of two of the "Preamble" (p) numbers before it
func xmas(xs []int, p int) int {
	for i := p; i < len(xs); i++ {
		ok := false
		for j := i - p; j < i; j++ {
			for k := i - p; k < i; k++ {
				if j != k {
					if xs[j]+xs[k] == xs[i] {
						ok = true
					}
				}
			}
		}

		if !ok {
			return xs[i]
		}
	}
	return 0
}

// Find the encryption weakness, add together the smallest and largest number
// in this contiguous range;
func weakness(xs []int, num int) int {
	for i, _ := range xs {
		ys := []int{}
		for j := i; j < len(xs); j++ {
			ys = append(ys, xs[j])
			if sum(ys) == num && len(ys) > 1 {
				sort.Ints(ys)
				return ys[0] + ys[len(ys)-1]
			}
		}
	}

	return 0
}

func sum(xs []int) int {
	s := 0
	for _, v := range xs {
		s = s + v
	}
	return s
}
