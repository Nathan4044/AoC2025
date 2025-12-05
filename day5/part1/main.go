package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type freshRange struct {
	start, end int
}

func main() {
	freshRanges, available := parseInput("input.txt")

	count := 0

	for _, ingredient := range available {
		for _, fresh := range freshRanges {
			if ingredient >= fresh.start && ingredient <= fresh.end {
				count++
				break
			}
		}
	}

	fmt.Println(count)
}

func parseInput(filename string) ([]freshRange, []int) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	ranges := []freshRange{}
	ingredients := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		index := strings.Index(line, "-")

		if index < 0 {
			break
		}

		start, err := strconv.Atoi(line[:index])
		if err != nil {
			panic(err)
		}

		end, err := strconv.Atoi(line[index+1:])
		if err != nil {
			panic(err)
		}

		ranges = append(
			ranges,
			freshRange{
				start: start,
				end:   end,
			},
		)
	}

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		ingredients = append(ingredients, num)
	}

	return ranges, ingredients
}
