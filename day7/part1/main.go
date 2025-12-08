package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

type set map[int]bool

func main() {
	view := parseInput("input.txt")

	beams := []int{bytes.Index(view[0], []byte{'S'})}
	splits := 0

	for _, line := range view[1:] {
		newBeams := set{}

		for _, n := range beams {
			switch line[n] {
			case '.':
				newBeams.add(n)
			case '^':
				newBeams.add(n-1, n+1)
				splits++
			default:
				panic("impossible byte")
			}
		}

		beams = newBeams.slice()
	}

	fmt.Println(splits)
}

func parseInput(filename string) [][]byte {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	lines := [][]byte{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, bytes.Clone(scanner.Bytes()))
	}

	return lines
}

func (s *set) add(vals ...int) {
	for _, v := range vals {
		(*s)[v] = true
	}
}

func (s *set) slice() []int {
	keys := []int{}

	for k := range *s {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	return keys
}
