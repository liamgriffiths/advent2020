package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const target = "shiny gold"

func main() {
	input := []string{}

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		input = append(input, sc.Text())
	}

	rules := []rule{}
	for _, x := range input {
		rule := parse(x)
		rules = append(rules, *rule)
	}

	res := findBagRules(rules, target)

	// lol, kind of a hacky way to get the unique values
	a := make(map[string]bool)
	for _, r := range res {
		k, _ := json.Marshal(r)
		a[string(k)] = true
	}

	fmt.Println(len(a))
}

func findBagRules(rules []rule, color string) []rule {
	found := []rule{}
	for _, r := range rules {
		for _, ch := range r.CanHave {
			if ch.Color == color {
				found = append(found, r)
			}
		}
	}

	for _, f := range found {
		results := findBagRules(rules, f.Color)
		found = append(found, results...)
	}

	return found
}

type rule struct {
	Color   string // containing bag color
	CanHave []canHave
}

type canHave struct {
	Number int    // max it can have
	Color  string // bag color
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

			ch := canHave{Number: number, Color: innerColor}
			canHaves = append(canHaves, ch)
		}
	}

	return &rule{
		Color:   color,
		CanHave: canHaves,
	}
}
