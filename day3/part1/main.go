package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	total := 0

	for scanner.Scan() {
		str := scanner.Text()
		line := []byte(str)

		firstPos, first := getByte(line[:len(line)-1])
		_, second := getByte(line[firstPos+1:])

		result := (int(first) - 48) * 10 + (int(second) - 48)
		total += result
	}

	fmt.Println(total)
}

func getByte(line []byte) (int, byte) {
	highest := line[0]
	highestPos := 0
	for i, c := range []byte(line) {
		if c > highest {
			highest = c
			highestPos = i
		}
	}

	return highestPos, highest
}
