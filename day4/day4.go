package main

import (
	"fmt"
	"os"
	"strings"
)

func parse() [][]int {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	rowsData := strings.Split(string(data), "\n")
	var values = make([][]int, len(rowsData))
	for i := range rowsData {
		var rowLength = len(rowsData[i])
		values[i] = make([]int, rowLength)
		for j := range rowLength {
			values[i][j] = 0
			if []rune(rowsData[i])[j] == '@' {
				values[i][j] = 1
			}
		}
	}

	return values
}

func print2DArray(value [][]int) {
	for _, row := range value {
		fmt.Println(row)
	}
}

func countSurrounding(grid [][]int, x, y int) int {
	var found = 0

	// ugh. ew
	if x > 0 && grid[y][x-1] == 1 {
		found += 1
	}
	if y > 0 && grid[y-1][x] == 1 {
		found += 1
	}
	if x < len(grid[0])-1 && grid[y][x+1] == 1 {
		found += 1
	}
	if y < len(grid)-1 && grid[y+1][x] == 1 {
		found += 1
	}
	if x > 0 && y > 0 && grid[y-1][x-1] == 1 {
		found += 1
	}
	if x > 0 && y < len(grid)-1 && grid[y+1][x-1] == 1 {
		found += 1
	}
	if x < len(grid[0])-1 && y > 0 && grid[y-1][x+1] == 1 {
		found += 1
	}
	if x < len(grid[0])-1 && y < len(grid)-1 && grid[y+1][x+1] == 1 {
		found += 1
	}

	return found
}

func findAccessibleRolls(grid [][]int, tooMany int) int {
	var accessible = 0

	for y := range len(grid) {
		for x := range len(grid[0]) {
			if grid[y][x] == 0 {
				continue
			}
			if countSurrounding(grid, x, y) < tooMany {
				fmt.Println("(", x, ",", y, ") is accessible")
				accessible += 1
			}
		}
	}

	return accessible
}

func main() {
	var grid = parse()
	print2DArray(grid)

	fmt.Println("Accessible rolls:", findAccessibleRolls(grid, 4))
}
