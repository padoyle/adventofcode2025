package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseBank(bankData string) []int {
	strValues := strings.Split(string(bankData), "")
	joltageValues := make([]int, len(strValues))

	for i := range len(strValues) {
		value, _ := strconv.Atoi(strValues[i])
		joltageValues[i] = value
	}

	return joltageValues
}

func parse() [][]int {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	bankData := strings.Split(string(data), "\n")

	banks := make([][]int, len(bankData))
	for i := range len(bankData) {
		banks[i] = parseBank(bankData[i])
	}

	return banks
}

func findBankValue(bank []int, batteryCount int) int {
	var result = make([]int, batteryCount)
	for i := range batteryCount {
		result[i] = 0
	}

	var p = 0
	for i := range bank {
		if bank[i] > result[p] {
			result[p] = bank[i]
		}

		// find leftmost place we can replace based on how many batteries remain
		var remainingValues = len(bank) - i
		var backtrack = max(batteryCount-remainingValues, 0)
		for q := backtrack; q < p; q++ {
			if bank[i] > result[q] {
				result[q] = bank[i]
				p = q
				// zero everything after our new position so we don't leave uncleared values
				for j := q + 1; j < batteryCount; j++ {
					result[j] = 0
				}
			}
		}
		// advance pointer by default; we'll backtrack in the next iteration if it's valid to do so
		if p < batteryCount-1 {
			p += 1
		}
	}

	var total, pow = 0, batteryCount - 1
	for i := range result {
		total += result[i] * int(math.Pow10(pow))
		pow -= 1
	}

	return total
}

func main() {
	var banks = parse()
	var total = 0
	// part 1
	for i := range len(banks) {
		var bank = banks[i]
		total += findBankValue(bank, 2)
	}

	fmt.Println("Part 1 total joltage:", total)

	// part 2
	total = 0
	for i := range len(banks) {
		var bank = banks[i]
		total += findBankValue(bank, 12)
	}
	fmt.Println("Part 2 total joltage:", total)
}
