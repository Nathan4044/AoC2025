package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	nodes := parseNodes("input.txt")
	fmt.Println(pathCount(nodes))
}

func pathCount(nodes map[string][]string) int {
	return traverse(nodes, "you", []string{})
}

func traverse(nodes map[string][]string, from string, traversed []string) int {
	reached := 0

	for _, node := range nodes[from] {
		if node == "out" {
			reached++
		} else if slices.Contains(traversed, node) {
			continue
		} else {
			reached += traverse(nodes, node, append(traversed, node))
		}
	}

	return reached
}

func parseNodes(filename string) map[string][]string {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	output := map[string][]string{}

	for scanner.Scan() {
		sections := strings.Split(scanner.Text(), ":")
		name := sections[0]
		outputs := strings.Fields(sections[1])

		output[name] = outputs
	}

	return output
}
