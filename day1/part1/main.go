package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("can't open file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	position := 50
	count := 0

	for scanner.Scan() {
		position += movement(scanner.Text())
		position = position % 100

		if position == 0 {
			count++
		}
	}
	
	fmt.Println(count)
}

func movement(line string) int {
	steps, err := strconv.Atoi(line[1:])
	if err != nil {
		panic("invalid number")
	}

	if line[0] == 'L' {
		steps = 100 - steps
	}

	return steps
}
