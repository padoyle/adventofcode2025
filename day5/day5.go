package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func parse() ([]Range, []int) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	strData := strings.Split(string(data), "\n")

	split := -1
	for i, value := range strData {
		if value == "" {
			split = i
			break
		}
	}
	if split == -1 {
		panic("couldn't find dividing line in data")
	}

	var freshRanges, ids = make([]Range, split), make([]int, len(strData)-split-1)

	for i := 0; i < split; i++ {
		rangeData := strings.Split(strData[i], "-")
		start, errStart := strconv.Atoi(rangeData[0])
		if errStart != nil {
			panic(errStart)
		}
		end, errEnd := strconv.Atoi(rangeData[1])
		if errEnd != nil {
			panic(errEnd)
		}

		freshRanges[i] = Range{start, end}
	}
	j := 0
	for i := split + 1; i < len(strData); i++ {
		id, err := strconv.Atoi(strData[i])
		if err != nil {
			panic(err)
		}
		ids[j] = id
		j += 1
	}

	return freshRanges, ids
}

func isInRange(id int, freshRanges []Range) bool {
	for _, r := range freshRanges {
		if id >= r.start && id <= r.end {
			return true
		}
	}
	return false
}

func countFreshIngredients(ids []int, freshRanges []Range) int {
	freshCount := 0
	for _, id := range ids {
		if isInRange(id, freshRanges) {
			freshCount += 1
		}
	}
	return freshCount
}

func condenseRanges(freshRangesSrc []Range) []Range {
	// copy for good measure
	freshRanges := make([]Range, len(freshRangesSrc))
	copy(freshRanges, freshRangesSrc)

	// 200 values at n^2 is not that bad
	for i, r1 := range freshRanges {
		for j := i + 1; j < len(freshRanges); j++ {
			r2 := freshRanges[j]
			if (r1.start <= r2.start && r1.end >= r2.start) || (r2.start <= r1.start && r2.end >= r1.start) {
				// clear current range and combine with later range
				// ideally I'd have a better way to "clear" this
				freshRanges[i] = Range{0, 0}
				freshRanges[j] = Range{min(r1.start, r2.start), max(r1.end, r2.end)}
				break
			}
		}
	}
	var freshRangesCondensed []Range
	for _, r := range freshRanges {
		if r.start == 0 && r.end == 0 {
			continue
		}
		freshRangesCondensed = append(freshRangesCondensed, r)
	}
	return freshRangesCondensed
}

func countAllPossibleFresh(freshRangesSrc []Range) int {
	rangesCondensed := condenseRanges(freshRangesSrc)

	total := 0
	for _, r := range rangesCondensed {
		total += (r.end - r.start) + 1
	}
	return total
}

func main() {
	var freshRanges, ids = parse()

	// part 1
	fmt.Println("Part 1 - Fresh ingredients:", countFreshIngredients(ids, freshRanges))
	// part 2
	fmt.Println("Part 2 - All possible fresh ingredients:", countAllPossibleFresh(freshRanges))
}
