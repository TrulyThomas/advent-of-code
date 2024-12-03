package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func read_input_bytes() (bytes []byte, err error) {
	fp, err := os.Open("input.txt")

	if err != nil {
		return
	}

	defer fp.Close()
	buf := make([]byte, 100)

	for {
		n, err := fp.Read(buf[0:])
		bytes = append(bytes, buf[0:n]...)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return bytes, err
}

func part1() {
	bytes, err := read_input_bytes()

	if err != nil {
		fmt.Print("whoptie, you suck")
		return
	}

	mul_re := regexp.MustCompile(`mul\([\d]{1,3},[\d]{1,3}\)`)
	mul_matches := mul_re.FindAll(bytes, -1)

	if mul_matches == nil {
		fmt.Println("No matches, something wrong")
		return
	}

	re := regexp.MustCompile(`\d+`)
	result := 0

	for _, input := range mul_matches {
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
