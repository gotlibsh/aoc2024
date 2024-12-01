package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func parseFile() ([]int, []int) {
	file, err := os.OpenFile("input.txt", os.O_RDONLY, os.ModePerm)
	checkError(err)
	defer file.Close()

	// group1 and group2
	group1 := make([]int, 10)
	group2 := make([]int, 10)

	// parse file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cols := strings.Split(scanner.Text(), "   ")

		num1, err := strconv.Atoi(cols[0])
		checkError(err)
		num2, err := strconv.Atoi(cols[1])
		checkError(err)

		group1 = append(group1, num1)
		group2 = append(group2, num2)
	}

	if len(group1) != len(group2) {
		panic("groups differ in size")
	}

	return group1, group2
}

func solve1() {
	group1, group2 := parseFile()

	// sort the groups in ascending order
	slices.Sort(group1)
	slices.Sort(group2)

	// calc diff
	var diff int64 = 0
	for i := range group1 {
		diff += int64(math.Abs(float64(group1[i] - group2[i])))
	}

	fmt.Printf("Part 1: %d\n", diff)
}

func solve2() {
	group1, group2 := parseFile()
	group2Counter := make(map[int]int)

	// count occurrences in group2
	for i := range group2 {
		group2Counter[group2[i]] += 1
	}

	// calc similarity
	similarity := 0
	for i := range group1 {
		similarity += (group1[i] * group2Counter[group1[i]])
	}

	fmt.Printf("Part 2: %d\n", similarity)
}

func main() {
	solve1()
	solve2()
}
