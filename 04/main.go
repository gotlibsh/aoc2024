package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func parseFile() []string {
	file, err := os.OpenFile("input.txt", os.O_RDONLY, os.ModePerm)
	checkError(err)
	defer file.Close()

	result := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result
}

func countXmas(puzzle []string) int {
	height, width := len(puzzle), len(puzzle[0])

	// left to right
	ltr := func(row int, col int) bool {
		return (width-col >= 3 &&
			puzzle[row][col] == 'M' &&
			puzzle[row][col+1] == 'A' &&
			puzzle[row][col+2] == 'S')
	}
	// right to left
	rtl := func(row int, col int) bool {
		return (col >= 2 &&
			puzzle[row][col] == 'M' &&
			puzzle[row][col-1] == 'A' &&
			puzzle[row][col-2] == 'S')
	}
	// top to bottom
	ttb := func(row int, col int) bool {
		return (height-row >= 3 &&
			puzzle[row][col] == 'M' &&
			puzzle[row+1][col] == 'A' &&
			puzzle[row+2][col] == 'S')
	}
	// bottom to top
	btt := func(row int, col int) bool {
		return (row >= 2 &&
			puzzle[row][col] == 'M' &&
			puzzle[row-1][col] == 'A' &&
			puzzle[row-2][col] == 'S')
	}
	// diagonal: left to right, top to bottom
	ltr_ttb := func(row int, col int) bool {
		return (width-col >= 3 && height-row >= 3 &&
			puzzle[row][col] == 'M' &&
			puzzle[row+1][col+1] == 'A' &&
			puzzle[row+2][col+2] == 'S')
	}
	// diagonal: left to right, bottom to top
	ltr_btt := func(row int, col int) bool {
		return (width-col >= 3 && row >= 2 &&
			puzzle[row][col] == 'M' &&
			puzzle[row-1][col+1] == 'A' &&
			puzzle[row-2][col+2] == 'S')
	}
	// diagonal: right to left, top to bottom
	rtl_ttb := func(row int, col int) bool {
		return (col >= 2 && height-row >= 3 &&
			puzzle[row][col] == 'M' &&
			puzzle[row+1][col-1] == 'A' &&
			puzzle[row+2][col-2] == 'S')
	}
	// diagonal: right to left, bottom to top
	rtl_btt := func(row int, col int) bool {
		return (col >= 2 && row >= 2 &&
			puzzle[row][col] == 'M' &&
			puzzle[row-1][col-1] == 'A' &&
			puzzle[row-2][col-2] == 'S')
	}
	bool_to_int := func(b bool) int {
		if b {
			return 1
		}
		return 0
	}

	count := 0
	for row, line := range puzzle {
		for col := range len(line) {
			if puzzle[row][col] != 'X' {
				continue
			}

			// check all 8 directions
			count += bool_to_int(ltr(row, col+1))
			count += bool_to_int(rtl(row, col-1))
			count += bool_to_int(ttb(row+1, col))
			count += bool_to_int(btt(row-1, col))
			count += bool_to_int(ltr_ttb(row+1, col+1))
			count += bool_to_int(ltr_btt(row-1, col+1))
			count += bool_to_int(rtl_ttb(row+1, col-1))
			count += bool_to_int(rtl_btt(row-1, col-1))
		}
	}

	return count
}

func solve1() {
	puzzle := parseFile()
	count := countXmas(puzzle)
	fmt.Printf("Part 1: %d\n", count)
}

func countX_mas(puzzle []string) int {
	height, width := len(puzzle), len(puzzle[0])

	// checker if (row, col) is in the middle of a X-MAS
	x_mas := func(row int, col int) bool {
		return ((puzzle[row-1][col-1] == 'M' && puzzle[row+1][col+1] == 'S') ||
			(puzzle[row-1][col-1] == 'S' && puzzle[row+1][col+1] == 'M')) &&
			((puzzle[row-1][col+1] == 'M' && puzzle[row+1][col-1] == 'S') ||
				(puzzle[row-1][col+1] == 'S' && puzzle[row+1][col-1] == 'M'))
	}

	count := 0
	for row := 1; row < height-1; row++ {
		for col := 1; col < width-1; col++ {
			if puzzle[row][col] == 'A' && x_mas(row, col) {
				count++
			}
		}
	}

	return count
}

func solve2() {
	puzzle := parseFile()
	count := countX_mas(puzzle)
	fmt.Printf("Part 2: %d\n", count)
}

func main() {
	solve1()
	solve2()
}
