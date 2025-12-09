package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func main() {
	points := parseInput("input.txt")

	area := 0

	for i, current := range points {
		for j := range i {
			newArea := areaOf(current, points[j])

			if newArea > area {
				area = newArea
			}
		}
	}

	fmt.Println(area)
}

func areaOf(a, b *point) int {
	width := (a.x - b.x)
	height := (a.y - b.y)

	if width < 0 {
		width *= -1
	}

	if height < 0 {
		height *= -1
	}

	width += 1
	height += 1

	return width * height
}

func parseInput(filename string) []*point {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	points := []*point{}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		numStrs := strings.Split(scanner.Text(), ",")

		p := &point{}

		num, err := strconv.Atoi(numStrs[0])
		if err != nil {
			panic(err)
		}

		p.x = num

		num, err = strconv.Atoi(numStrs[1])
		if err != nil {
			panic(err)
		}

		p.y = num

		points = append(points, p)
	}

	return points
}
