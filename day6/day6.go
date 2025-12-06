package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const OP_ADD = '+'
const OP_MULT = '*'

type Problem struct {
	operands  []int
	operation byte
}

func parse() []Problem {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	strData := strings.Split(string(data), "\n")

	var problems []Problem

	// parse this down into something readable
	for line := range strData {
		p := 0
		lineData := strings.Split(strData[line], " ")
		for _, lineValue := range lineData {
			if lineValue == "" {
				continue
			} else if lineValue == "*" || lineValue == "+" {
				if len(problems) <= p {
					problems = append(problems, Problem{})
				}
				problems[p].operation = lineValue[0]
			} else {
				value, err := strconv.Atoi(lineValue)
				if err != nil {
					panic(err)
				}
				if len(problems) <= p {
					problems = append(problems, Problem{})
				}
				problems[p].operands = append(problems[p].operands, value)
			}
			p += 1
		}
	}

	return problems
}

// guess we need a damn grid
func parsePart2() []Problem {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	strData := strings.Split(string(data), "\n")

	var problems []Problem
	var runes = make([][]rune, len(strData[0]))

	// lets parse this into an array of columns, for the sake of my sanity
	for _, line := range strData {
		for i, b := range line {
			runes[i] = append(runes[i], b)
		}
	}

	// add and mult are both commutative, so parse order is blessedly irrelevant
	p := Problem{}
	for colId, colValue := range runes {
		wasEmpty, power, value := true, 0, 0
		for rowId := len(colValue) - 1; rowId >= 0; rowId-- {
			rowVal := runes[colId][rowId]
			if rowVal == OP_ADD || rowVal == OP_MULT {
				p.operation = byte(rowVal)
			} else if rowVal != ' ' {
				wasEmpty = false
				value += int(rowVal-'0') * int(math.Pow10(power))
				power += 1
			}
		}
		if value != 0 {
			p.operands = append(p.operands, value)
		}
		// finished parsing a problem when we see a fully empty column
		if wasEmpty || colId == len(runes)-1 {
			problems = append(problems, p)
			p = Problem{}
		}
	}

	return problems
}

func doMath(problems []Problem) int {
	total := 0
	for _, problem := range problems {
		if problem.operation == OP_ADD {
			value := 0
			for _, operand := range problem.operands {
				value += operand
			}
			total += value
		}
		if problem.operation == OP_MULT {
			value := 1
			for _, operand := range problem.operands {
				value *= operand
			}
			total += value
		}
	}

	return total
}

func main() {
	problems := parse()
	fmt.Println("part 1:", doMath((problems)))

	problems2 := parsePart2()
	fmt.Println("part 2:", doMath((problems2)))
}
