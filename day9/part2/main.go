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

type line interface {
	intersects(l line) bool
}

type horizontal struct {
	y            int
	xStart, xEnd int
}

func (h *horizontal) intersects(l line) bool {
	v, ok := l.(*vertical)
	if ok {
		if v.x <= h.xStart || v.x >= h.xEnd {
			return false
		}

		if h.y <= v.yStart || h.y >= v.yEnd {
			return false
		}

		return true
	}

	o, _ := l.(*horizontal)

	if o.y != h.y {
		return false
	}

	if o.xEnd <= h.xStart || o.xStart >= h.xEnd {
		return false
	}

	if o.xStart >= h.xStart && o.xStart < h.xEnd && o.xEnd <= h.xEnd {
		return false
	}

	return true
}

type vertical struct {
	x            int
	yStart, yEnd int
}

func (v *vertical) intersects(l line) bool {
	h, ok := l.(*horizontal)
	if ok {
		return h.intersects(v)
	}

	o, _ := l.(*vertical)

	if o.x != v.x {
		return false
	}

	if o.yEnd <= v.yStart || o.yStart >= v.yEnd {
		return false
	}

	if o.yStart >= v.yStart && o.yStart < v.yEnd && o.yEnd <= v.yEnd {
		return false
	}

	return true
}

func main() {
	points := parseInput("input.txt")
	lines := makeLines(points)

	fmt.Println(largestArea(points, lines))
}

func makeLines(points []*point) []line {
	lines := []line{}
	for i := range len(points) - 1 {
		lines = append(lines, points[i].to(points[i+1]))
	}

	lines = append(lines, points[len(points)-1].to(points[0]))

	return lines
}

func largestArea(points []*point, lines []line) int {
	area := 0

	for i, current := range points {
		for j := range i {
			newArea := areaOf(current, points[j])

			if newArea > area && validRect(current, points[j], lines) {
				area = newArea
			}
		}
	}

	return area
}

func validRect(a, b *point, lines []line) bool {
	var rectLines []line

	if a.inLineWith(b) {
		rectLines = []line{
			a.to(b),
		}
	} else {
		c := &point{
			x: a.x,
			y: b.y,
		}
		d := &point{
			x: b.x,
			y: a.y,
		}

		rectLines = []line{
			a.to(c),
			c.to(b),
			b.to(d),
			d.to(a),
		}
	}

	for _, l := range lines {
		for _, r := range rectLines {
			if l.intersects(r) {
				return false
			}
		}
	}

	return true
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

func (p *point) to(o *point) line {
	var l line

	if p.x == o.x {
		l = &vertical{
			x:      p.x,
			yStart: min(p.y, o.y),
			yEnd:   max(p.y, o.y),
		}
	} else if p.y == o.y {
		l = &horizontal{
			y:      p.y,
			xStart: min(p.x, o.x),
			xEnd:   max(p.x, o.x),
		}
	}

	return l
}

func (p *point) inLineWith(o *point) bool {
	return p.x == o.x || p.y == o.y
}
