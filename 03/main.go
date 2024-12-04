package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func parseFile() string {
	data, err := os.ReadFile("input.txt")
	checkError(err)

	return string(data)
}

func sumMuls(data string) int {
	mulRe := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := mulRe.FindAllStringSubmatch(data, -1)

	sum := 0
	for _, match := range matches {
		op1, _ := strconv.Atoi(string(match[1]))
		op2, _ := strconv.Atoi(string(match[2]))
		sum += op1 * op2
	}

	return sum
}

func solve1() {
	data := parseFile()
	sum := sumMuls(data)
	fmt.Printf("Part 1: %d\n", sum)
}

func getEnabledPart(data string) string {
	var result strings.Builder
	var loc []int

	doRe := regexp.MustCompile(`do\(\)`)
	dontRe := regexp.MustCompile(`don't\(\)`)

	index := 0
	for index < len(data) {
		remain := data[index:]

		// find next don't()
		loc = dontRe.FindStringIndex(remain)
		if loc == nil {
			// no more don't()s
			result.WriteString(remain)
			break
		}

		// store enabled part
		result.WriteString(remain[:loc[0]])
		index += loc[1]

		// find next do()
		remain = data[index:]
		loc = doRe.FindStringIndex(remain)
		if loc == nil {
			// found don't() with no followup do()
			break
		}
		index += loc[1]
	}

	return result.String()
}

func solve2() {
	data := parseFile()
	enabled := getEnabledPart(data)
	sum := sumMuls(enabled)
	fmt.Printf("Part 2: %d\n", sum)
}

func main() {
	solve1()
	solve2()
}
