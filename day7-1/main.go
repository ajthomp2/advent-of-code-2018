package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
	"time"
)

type singleLinkedNode struct {
	id string
	prev []*singleLinkedNode
	start bool
}

func main() {
	start := time.Now()

	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal("error reading data")
	}

	instructions := strings.Split(string(data), "\r\n")

	nodes := make(map[string]*singleLinkedNode)
	for _, instruction := range instructions {
		words := strings.Split(instruction, " ")

		start, end := words[1], words[7]
		startNode, ok := nodes[start]
		if !ok {
			startNode = &singleLinkedNode{id: start, start: true}
			nodes[start] = startNode
		}
		endNode, ok := nodes[end]
		if !ok {
			endNode = &singleLinkedNode{id: end}
			nodes[end] = endNode
		}
		endNode.start = false

		endNode.prev = append(endNode.prev, startNode)
	}

	var nextNodes []string
	for k, v := range nodes {
		if len(v.prev) == 0 {
			nextNodes = append(nextNodes, k)
			delete(nodes, k)
		}
	}

	var order []string
	for len(nextNodes) != 0 {
		sort.Strings(nextNodes)
		n := nextNodes[0]
		order = append(order, n)
		nextNodes = nextNodes[1:]

		var nextLevel []string
		for k, v := range nodes {
			for i := range v.prev {
				if v.prev[i].id == n {
					v.prev = append(v.prev[:i], v.prev[i+1:]...)
					break
				}
			}
			if len(v.prev) == 0 {
				nextLevel = append(nextLevel, k)
				delete(nodes, k)
			}
		}
		if len(nextLevel) != 0 {
			sort.Strings(nextLevel)
			nextNodes = append(nextNodes, nextLevel...)
		}
	}

	fmt.Println(strings.Join(order, ""))
	fmt.Println("\n", time.Since(start))
}
