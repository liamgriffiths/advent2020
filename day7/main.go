package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := []string{}

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		input = append(input, sc.Text())
	}

	for _, x := range input {
		fmt.Println(parse(x))
	}

}

type rule struct {
	color   string // containing bag color
	canHave []canHave
}

type canHave struct {
	number int    // max it can have
	color  string // bag color
}

func parse(input string) *rule {
	re := regexp.MustCompile(`^(?P<color>.*) bags contain (?P<contains>.*)\.$`)
	names := re.SubexpNames()
	m := re.FindAllStringSubmatch(input, -1)[0]

	kv := map[string]string{}
	for i, n := range m {
		kv[names[i]] = n
	}

	color := strings.Trim(kv["color"], " ")
	canHaves := []canHave{}

	if kv["contains"] != "no other bags" {
		bs := strings.Split(kv["contains"], ",")
		for _, b := range bs {
			b = strings.Trim(b, " ")
			re := regexp.MustCompile(`^(?P<count>\d+) (?P<color>.*) bag`)
			names := re.SubexpNames()
			m := re.FindAllStringSubmatch(b, -1)[0]

			kv := map[string]string{}
			for i, n := range m {
				kv[names[i]] = n
			}

			number, _ := strconv.Atoi(kv["count"])
			innerColor := strings.Trim(kv["color"], " ")

			canHaves = append(canHaves, canHave{number: number, color: innerColor})
		}
	}

	return &rule{
		color:   color,
		canHave: canHaves,
	}
}
