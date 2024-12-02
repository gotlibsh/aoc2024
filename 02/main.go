package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func lineToReport(line string) []int {
	levels := strings.Split(line, " ")
	report := make([]int, len(levels))

	for i, level := range levels {
		x, err := strconv.Atoi(level)
		checkError(err)
		report[i] = x
	}

	return report
}

func parseFile() [][]int {
	file, err := os.OpenFile("input.txt", os.O_RDONLY, os.ModePerm)
	checkError(err)
	defer file.Close()

	// collect reports
	reports := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		report := lineToReport(scanner.Text())
		reports = append(reports, report)
	}

	return reports
}

func isSafeReport(report []int) bool {
	// record inc/dec order
	sign := report[1] - report[0]

	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]

		// is distance from neighbor ok
		if diff < -3 || diff == 0 || diff > 3 || diff*sign < 0 {
			return false
		}
	}

	return true
}

func solve1() {
	reports := parseFile()
	counter := 0

	for _, report := range reports {
		if isSafeReport(report) {
			counter++
		}
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func isSafeReportProblemDampener(report []int) bool {
	concatSlices := func(s1 []int, s2 []int) []int {
		concat := make([]int, len(s1))
		copy(concat, s1)
		return append(concat, s2...)
	}

	for i := 0; i < len(report); i++ {
		// check if safe when discarding level i
		reducedReport := concatSlices(report[:i], report[i+1:])

		if isSafeReport(reducedReport) {
			return true
		}
	}

	return false
}

func solve2() {
	reports := parseFile()
	counter := 0

	for _, report := range reports {
		if isSafeReportProblemDampener(report) {
			counter++
		}
	}

	fmt.Printf("Part 2: %d\n", counter)
}

func main() {
	solve1()
	solve2()
}
