package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type floorMap struct {
	rows          [][]byte
	height, width int
}

func main() {
	floor := parseMap("input.txt")

	total := markReachable(floor)

	fmt.Println(total)
}

func markReachable(floor floorMap) int {
	count := 0
	for y, row := range floor.rows {
		for x, char := range row {
			if char == '@' {
				if countNeighbours(floor, x, y) < 4 {
					floor.rows[y][x] = 'x'
					count++
				}
			}
		}
	}

	return count
}

func countNeighbours(floor floorMap, x, y int) int {
	count := 0

	for j := -1; j <= 1; j++ {
		if y+j < 0 || y+j >= floor.height {
			continue
		}
		for i := -1; i <= 1; i++ {
			if x+i < 0 || x+i >= floor.width {
				continue
			}

			if i == 0 && j == 0 {
				continue
			}

			if floor.rows[y+j][x+i] == '@' || floor.rows[y+j][x+i] == 'x' {
				count++
			}
		}
	}

	return count
}

func parseMap(filename string) floorMap {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	floor := floorMap{
		rows: [][]byte{},
	}

	for scanner.Scan() {
		floor.rows = append(floor.rows, bytes.Clone(scanner.Bytes()))
	}

	floor.height = len(floor.rows)
	floor.width = len(floor.rows[0])

	return floor
}
