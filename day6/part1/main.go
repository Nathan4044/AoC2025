package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := parseInput("input.txt")

	total := 0
	for i := 0; i < len(lines[0]); i++ {
		columnTotal := 0
		switch lines[len(lines)-1][i] {
		case "+":
			for j := 0; j < len(lines)-1; j++ {
				num, err := strconv.Atoi(lines[j][i])
				if err != nil {
					panic(err)
				}
				columnTotal += num
			}
		case "*":
			columnTotal = 1
			for j := 0; j < len(lines)-1; j++ {
				num, err := strconv.Atoi(lines[j][i])
				if err != nil {
					panic(err)
				}
				columnTotal *= num
			}
		default:
			panic("impossible operator")
		}
		total += columnTotal
	}

	fmt.Println(total)
}

func parseInput(filename string) [][]string {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	lines := [][]string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		sections := strings.Fields(scanner.Text())
		lines = append(lines, sections)
	}

	return lines
}
