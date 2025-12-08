package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type problem struct {
	nums []int
	add  bool
}

func main() {
	rotated := rotateInput("input.txt")
	problems := parseProblems(rotated)

	total := 0

	for _, p := range problems {
		cur := 0
		if p.add {
			for _, n := range p.nums {
				cur += n
			}
		} else {
			cur = 1
			for _, n := range p.nums {
				cur *= n
			}
		}
		total += cur
	}

	fmt.Println(total)
}

func parseProblems(lines []string) []problem {
	problems := []problem{}
	current := problem{nums: []int{}}

	for _, line := range lines {
		s := strings.Trim(line[:len(line)-1], " ")
		if s == "" {
			problems = append(problems, current)
			current = problem{nums: []int{}}
		} else {
			if line[len(line)-1] == '+' {
				current.add = true
			}
			num, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			current.nums = append(current.nums, num)
		}
	}

	problems = append(problems, current)

	return problems
}

func rotateInput(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	lines := [][]byte{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		for i, b := range scanner.Bytes() {
			if len(lines) <= i {
				lines = append(lines, []byte{})
			}
			lines[i] = append(lines[i], b)
		}
	}

	strs := []string{}
	for _, line := range lines {
		strs = append(strs, string(line))
	}

	return strs
}
