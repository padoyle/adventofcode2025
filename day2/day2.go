package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var factorsMap = map[int][]int{
	2:  []int{2},
	3:  []int{3},
	4:  []int{2, 4},
	5:  []int{5},
	6:  []int{2, 3, 6},
	7:  []int{7},
	8:  []int{2, 4, 8},
	9:  []int{3, 9},
	10: []int{2, 5, 10},
	11: []int{11},
	12: []int{2, 3, 4, 6, 12},
}

type Range struct {
	first int
	last  int
}

func parseRange(rangeData string) Range {
	values := strings.Split(string(rangeData), "-")
	first, _ := strconv.Atoi(values[0])
	last, _ := strconv.Atoi(values[1])

	return Range{first, last}
}

func parse() []Range {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	rangeData := strings.Split(string(data), ",")

	ranges := make([]Range, len(rangeData))
	for i := range len(rangeData) {
		ranges[i] = parseRange(rangeData[i])
	}

	return ranges
}

func isValid(id int) bool {
	var digits = int(math.Log10(float64(id))) + 1
	if digits%2 == 1 {
		return true
	}
	pow := int(math.Pow10((digits / 2)))
	upper := id / pow
	lower := id - (upper * pow)

	return upper != lower
}

func isValidPart2(id int) bool {
	var idStr = strconv.Itoa(id)

	var factors = factorsMap[len(idStr)]

	for i := range factors {
		var l = len(idStr) / factors[i]
		var firstSlice = idStr[:l]
		var sliceRepeats = true
		for j := l; j < len(idStr); j += l {
			if idStr[j:j+l] != firstSlice {
				sliceRepeats = false
				break
			}
		}
		if sliceRepeats {
			fmt.Println(idStr, "is invalid:", firstSlice, "repeats")
			return false
		}
	}

	return true
}

func totalInvalidIds(r Range) int {
	var invalidTotal = 0
	for offset := range (r.last - r.first) + 1 {
		var id = r.first + offset
		if !isValidPart2(id) {
			invalidTotal += id
		}
	}
	fmt.Println("invalid ids in range total to:", invalidTotal)
	return invalidTotal
}

func main() {
	var ranges = parse()
	var totalInvalid = 0
	for i := range ranges {
		var numRange = ranges[i]
		fmt.Println("range: ", numRange.first, "-", numRange.last)
		totalInvalid += totalInvalidIds(numRange)
		fmt.Println()
	}
	fmt.Println("Total invalid ids:", totalInvalid)
}
