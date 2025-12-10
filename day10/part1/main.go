package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type machine struct {
	lights      uint
	lightLength uint
	buttons     []uint
	joltage     []uint
}

func main() {
	machines := parseMachines("input.txt")

	var result uint = 0
	for _, m := range machines {
		result += solveMachine(0, &m, 20)
	}

	fmt.Println(result)
}

func solveMachine(state uint, m *machine, limit uint) uint {
	if limit == 0 {
		return math.MaxUint
	}

	if state == m.lights {
		return 0
	}

	var lowest uint = math.MaxUint
	for _, button := range m.buttons {
		newState := state ^ button
		depth := solveMachine(newState, m, min(limit-1, lowest))

		lowest = min(lowest, depth)
	}

	if lowest == math.MaxUint {
		return math.MaxUint
	} else {
		return lowest + 1
	}
}

func parseMachines(filename string) []machine {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	machines := []machine{}

	for scanner.Scan() {
		machines = append(machines, parseMachine(scanner.Text()))
	}

	return machines
}

func parseMachine(line string) machine {
	sections := strings.Fields(line)

	lights, lightLength := parseLight(sections[0])

	return machine{
		lights:      lights,
		lightLength: lightLength,
		buttons:     parseButtons(sections[1:len(sections)-1], lightLength),
		joltage:     parseJoltage(sections[len(sections)-1]),
	}
}

func parseLight(str string) (uint, uint) {
	var result uint = 0
	var position uint = 1
	var count uint = 0
	for i := len(str) - 2; i > 0; i-- {
		if str[i] == '#' {
			result += position
		}

		count++
		position *= 2
	}

	return result, count
}

func parseButtons(strs []string, lightLength uint) []uint {
	result := make([]uint, 0, len(strs))

	for _, str := range strs {
		numStrs := strings.Split(str[1:len(str)-1], ",")
		var num uint = 0

		for _, s := range numStrs {
			n, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}

			num += uint(math.Exp2(float64(lightLength - uint(n) - 1)))
		}

		result = append(result, num)
	}
	return result
}

func parseJoltage(str string) []uint {
	numStrs := strings.Split(str[1:len(str)-1], ",")
	nums := make([]uint, 0, len(numStrs))

	for _, s := range numStrs {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		nums = append(nums, uint(n))
	}

	return nums
}
