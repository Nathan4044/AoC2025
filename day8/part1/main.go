package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type box struct {
	x, y, z int
	chain   int
}

type distance struct {
	left, right *box
	distance    float64
}

func main() {
	boxes, distances := parseBoxes("input.txt")
	constructChains(distances, boxes, 1000)

	allLengths := map[int]int{}
	for _, b := range boxes {
		allLengths[b.chain]++
	}

	chainLengths := []int{}
	for k, v := range allLengths {
		if k != 0 {
			chainLengths = append(chainLengths, v)
		}
	}
	slices.Sort(chainLengths)
	slices.Reverse(chainLengths)

	fmt.Println(chainLengths[0] * chainLengths[1] * chainLengths[2])
}

func constructChains(distances []distance, boxes []*box, chainLimit int) {
	currentChain := 1

	for i := range chainLimit {
		left := distances[i].left
		right := distances[i].right

		if left.isConnected(right) {
			continue
		}

		if left.chain > 0 && right.chain > 0 {
			mergingChain := right.chain

			for _, b := range boxes {
				if b.chain == mergingChain {
					b.chain = left.chain
				}
			}
		}

		existingChain := max(left.chain, right.chain)

		if existingChain == 0 {
			existingChain = currentChain
			currentChain++
		}

		left.chain = existingChain
		right.chain = existingChain
	}
}

func parseBoxes(filename string) ([]*box, []distance) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	boxes := []*box{}
	distances := []distance{}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		numStrings := strings.Split(scanner.Text(), ",")
		nums := make([]int, 3)

		for i, s := range numStrings {
			n, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			nums[i] = n
		}

		newBox := &box{
			x: nums[0],
			y: nums[1],
			z: nums[2],
		}

		for _, b := range boxes {
			distances = append(distances, b.distanceFrom(newBox))
		}

		boxes = append(boxes, newBox)
	}

	slices.SortFunc(distances, func(a, b distance) int {
		if a.distance < b.distance {
			return -1
		}
		if a.distance > b.distance {
			return 1
		}
		return 0
	})

	return boxes, distances
}

func (b *box) distanceFrom(o *box) distance {
	dist := math.Sqrt(
		(math.Abs(float64(b.x-o.x)) * math.Abs(float64(b.x-o.x))) +
			(math.Abs(float64(b.y-o.y)) * math.Abs(float64(b.y-o.y))) +
			(math.Abs(float64(b.z-o.z)) * math.Abs(float64(b.z-o.z))),
	)

	return distance{
		left:     b,
		right:    o,
		distance: dist,
	}
}

func (b *box) isConnected(o *box) bool {
	return b.chain != 0 && b.chain == o.chain
}

func (b *box) String() string {
	return fmt.Sprintf("%d, %d, %d (%d)", b.x, b.y, b.z, b.chain)
}
