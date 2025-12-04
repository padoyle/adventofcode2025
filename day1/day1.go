package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse() []string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n")
}

func parse_line(line string) int {
	rotation, err := strconv.Atoi(line[1:])
	if err != nil {
		fmt.Println("ERROR:", err)
		panic(err)
	}
	if line[0] == 'L' {
		rotation *= -1
	}
	return rotation
}

// this was double counting... something, somewhere, so I decided to go crude
// func rotate(start, value int) (int, int) {
// 	var zeroes = 0
// 	var result = start + value
// 	for result < 0 {
// 		result += 100
// 		zeroes += 1
// 	}
// 	for result > 99 {
// 		result -= 100
// 		zeroes += 1
// 	}
// 	if result == 0 && (value > 100 || value < 0) {
// 		zeroes += 1
// 	}
// 	if start == 0 && value < 0 {
// 		zeroes -= 1
// 	}

// 	return result, zeroes
// }

func rotate_basic(start, value int) int {
	var result = start + value
	if result < 0 {
		result += 100
	}
	if result > 99 {
		result -= 100
	}
	return result
}

func rotate_inc(start, value int) (int, int) {
	var result, zeroes = start, 0
	var increment = 1
	if value < 0 {
		value *= -1
		increment = -1
	}
	for range value {
		result = rotate_basic(result, increment)
		if result == 0 {
			zeroes += 1
		}
	}

	return result, zeroes
}

func main() {
	var input = parse()
	var current = 50
	var count = 0
	for i := range input {
		new_value, times_crossed_zero := rotate_inc(current, parse_line(input[i]))

		current = new_value
		count += times_crossed_zero
	}

	fmt.Println("Hit 0", count, "times")
}
