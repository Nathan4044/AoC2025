package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	total := 0

	for scanner.Scan() {
		str := scanner.Text()
		total += calcJoltage([]byte(str))
	}

	fmt.Println(total)
}

func calcJoltage(line []byte) int {
	digits := []byte{}
	startPos := 0

	for i := 11; i >= 0; i-- {
		newPos, digit := getByte(line[startPos:len(line)-i])
		digits = append(digits, digit)
		startPos += newPos + 1
	}

	result, err := strconv.Atoi(string(digits))
	if err != nil {
		panic(err)
	}

	return result
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
