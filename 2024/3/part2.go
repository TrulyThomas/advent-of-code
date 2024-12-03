package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func part2() {
	bytes, err := read_input_bytes()

	if err != nil {
		fmt.Print("whoptie, you suck")
		return
	}

	mul_re := regexp.MustCompile(`mul\([\d]{1,3},[\d]{1,3}\)|do\(\)|don\'t\(\)`)
	mul_matches := mul_re.FindAll(bytes, -1)

	if mul_matches == nil {
		fmt.Println("No matches, something wrong")
		return
	}

	re := regexp.MustCompile(`\d+`)
	do_re := regexp.MustCompile(`do\(\)`)
	dont_re := regexp.MustCompile(`don\'t\(\)`)

	skip := false
	result := 0

	for _, input := range mul_matches {
		do_match := do_re.Find(input)

		if do_match != nil {
			skip = false
			continue
		}

		dont_match := dont_re.Find(input)

		if dont_match != nil {
			skip = true
			continue
		}

		if skip {
			continue
		}

		matches := re.FindAll(input, -1)

		if len(matches) < 2 {
			fmt.Println("Not enough matches to form a tuple, something is wrong")
			return
		}

		num1, err1 := strconv.Atoi(string(matches[0]))
		num2, err2 := strconv.Atoi(string(matches[1]))
		if err1 != nil || err2 != nil {
			fmt.Println("Error converting bytes to strings to integers")
			return
		}
		result += num1 * num2
	}
	fmt.Println("Result: ", result)
}
