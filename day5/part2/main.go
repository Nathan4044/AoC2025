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
	freshRanges := parseInput("input.txt")
	reduced := true

	for reduced {
		freshRanges, reduced = consolidateRanges(freshRanges)
	}

	total := 0

	for _, fresh := range freshRanges {
		count := fresh.end - fresh.start + 1
		total += count
	}

	fmt.Println(total)
}

func consolidateRanges(original []freshRange) ([]freshRange, bool) {
	result := []freshRange{}
	reduced := false

	for _, orig := range original {
		absorbed := false
		for i, current := range result {
			if current.contains(orig.start) || orig.contains(current.start) {
				result[i].combine(orig)
				reduced = true
				absorbed = true
				break
			}
		}

		if !absorbed {
			result = append(result, orig)
		}
	}

	return result, reduced
}

func (fr *freshRange) contains(val int) bool {
	return val >= fr.start && val <= fr.end
}

func (fr *freshRange) combine(other freshRange) {
	fr.start = min(fr.start, other.start)
	fr.end = max(fr.end, other.end)
}

func parseInput(filename string) []freshRange {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	ranges := []freshRange{}

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

	return ranges
}
