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
		move := movement(scanner.Text())

		for move > 99 {
			count++
			move -= 100
		}

		for move < -99 {
			count++
			move += 100
		}

		newPos := position + move

		if newPos < 0 {
			if position != 0 {
				count++
			}

			newPos += 100
		} else if newPos == 0 {
			count++
		} else if newPos > 99 {
			newPos -= 100
			count++
		}
		
		position = newPos
	}

	fmt.Println(count)
}

func movement(line string) int {
	steps, err := strconv.Atoi(line[1:])
	if err != nil {
		panic("invalid number")
	}

	if line[0] == 'L' {
		steps *= -1
	}

	return steps
}
