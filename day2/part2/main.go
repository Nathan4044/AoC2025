package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	ranges := parseRanges("input.txt")
	// ranges := parseRanges("test_input.txt")

	count := 0
	for _, r := range ranges {
		count += findInvalid(r)
	}

	fmt.Println(count)
}

type pair struct {
	start int
	end   int
}

func parseRanges(filename string) []pair {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	data, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	entries := strings.Split(string(data[:len(data)-1]), ",")

	pairs := make([]pair, 0, len(entries))

	for _, e := range entries {
		nums := strings.Split(e, "-")

		if len(nums) != 2 {
			panic(e)
		}

		p := pair{}

		n, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}

		p.start = n

		n, err = strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}

		p.end = n

		pairs = append(pairs, p)
	}

	return pairs
}

func findInvalid(r pair) int {
	total := 0

	for i := r.start; i <= r.end; i++ {
		str := strconv.Itoa(i)

		if testInvalid(str) {
			total += i
		}
	}

	return total
}

func testInvalid(str string) bool {
	for j := 2; j <= len(str); j++ {
		if len(str) % j != 0 {
			continue
		}

		testStr := strings.Repeat(str[:len(str)/j], j)

		if strings.Compare(str, testStr) == 0 {
			return true
		}
	}

	return false
}
