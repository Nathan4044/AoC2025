package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type node struct {
	val   string
	paths []*node
	count int
	gates int
}

func main() {
	nodeMap := parseNodes("input.txt")
	nodes := &map[string]*node{}
	(*nodes)["out"] = &node{val: "out", paths: []*node{}, count: -1}
	tree := paths(nodeMap, nodes, "svr")

	count, _ := traverse(tree, []*node{})

	fmt.Println(count)
}

func paths(nodes map[string][]string, existing *map[string]*node, name string) *node {
	n := &node{
		val:   name,
		paths: []*node{},
		count: -1,
	}

	(*existing)[name] = n

	for _, next := range nodes[name] {
		e, ok := (*existing)[next]
		if ok {
			n.paths = append(n.paths, e)
		} else {
			n.paths = append(n.paths, paths(nodes, existing, next))
		}
	}

	return n
}

func traverse(n *node, visited []*node) (int, int) {
	if n.count >= 0 {
		return n.count, n.gates
	}

	if n.val == "out" {
		return 1, 0
	}

	newVisited := append(visited, n)

	count := 0
	gates := 0

	for _, next := range n.paths {
		c, g := traverse(next, newVisited)

		if g > gates {
			gates = g
			count = c
		} else if g == gates {
			count += c
		}
	}

	if n.val == "fft" || n.val == "dac" {
		gates++
	}

	n.count = count
	n.gates = gates

	return count, gates
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
