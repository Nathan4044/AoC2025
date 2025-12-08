package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	view := parseInput("example_input.txt")

	beams := map[int]int{}
	beams[bytes.Index(view[0], []byte{'S'})]++

	for _, line := range view[1:] {
		newBeams := map[int]int{}

		for k, v := range beams {
			switch line[k] {
			case '.':
				newBeams[k] += v
			case '^':
				newBeams[k-1] += v
				newBeams[k+1] += v
			default:
				panic("impossible byte")
			}
		}

		beams = newBeams
	}

	total := 0
	for _, v := range beams {
		total += v
	}

	fmt.Println(total)
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
