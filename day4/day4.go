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

func print2DArray(grid [][]int) {
	for _, row := range grid {
		for _, value := range row {
			if value == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("@")
			}
		}
		fmt.Println()
	}
	fmt.Println()
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

	for y := range grid {
		for x := range y {
			if grid[y][x] == 0 {
				continue
			}
			if countSurrounding(grid, x, y) < tooMany {
				accessible += 1
			}
		}
	}

	return accessible
}

func removeAccessibleRolls(grid [][]int, tooMany int) ([][]int, int) {
	var accessible = 0
	var newGrid = make([][]int, len(grid))

	for y := range len(grid) {
		newGrid[y] = make([]int, len(grid[y]))
		for x := range len(grid[0]) {
			if grid[y][x] == 0 {
				newGrid[y][x] = 0
				continue
			}
			if countSurrounding(grid, x, y) < tooMany {
				newGrid[y][x] = 0
				accessible += 1
			} else {
				newGrid[y][x] = 1
			}
		}
	}

	return newGrid, accessible
}

func removeAllPossible(grid [][]int, tooMany int) int {
	var total = 0
	var workingGrid, removed = grid, 0
	for {
		workingGrid, removed = removeAccessibleRolls(workingGrid, tooMany)
		fmt.Println("Removed", removed)
		// print2DArray(workingGrid)

		total += removed

		if removed == 0 {
			break
		}
	}

	return total
}

func main() {
	var grid = parse()
	fmt.Print("Part 1:\n")
	print2DArray(grid)

	// part 1
	fmt.Println("Accessible rolls:", findAccessibleRolls(grid, 4))

	// part 2
	fmt.Print("\n\nPart 2:\n")
	fmt.Println(removeAllPossible(grid, 4))
}
