package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
	"time"
)

type doubleLinkedNode struct {
	id string
	next []*doubleLinkedNode
	prev []*doubleLinkedNode
}

type sortNodes []*doubleLinkedNode

func main() {
	start := time.Now()

	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal("error reading data")
	}

	instructions := strings.Split(string(data), "\r\n")

	nodes := make(map[string]*doubleLinkedNode)
	for _, instruction := range instructions {
		words := strings.Split(instruction, " ")

		start, end := words[1], words[7]
		startNode, ok := nodes[start]
		if !ok {
			startNode = &doubleLinkedNode{id: start}
			nodes[start] = startNode
		}
		endNode, ok := nodes[end]
		if !ok {
			endNode = &doubleLinkedNode{id: end}
			nodes[end] = endNode
		}

		startNode.next = append(startNode.next, endNode)
		endNode.prev = append(endNode.prev, startNode)
	}

	// get starting nodes
	var nextNodes []*doubleLinkedNode
	for _, v := range nodes {
		if len(v.prev) == 0 {
			nextNodes = append(nextNodes, v)
		}
	}

	var order bytes.Buffer
	for len(nextNodes) != 0 {
		sort.Sort(sortNodes(nextNodes))
		n := nextNodes[0]
		nextNodes = nextNodes[1:]

		for _, child := range n.next {
			for i := range child.prev {
				if child.prev[i].id == n.id {
					child.prev = append(child.prev[:i], child.prev[i+1:]...)
					break
				}
			}
			if len(child.prev) == 0 {
				nextNodes = append(nextNodes, child)
			}
		}

		order.WriteString(n.id)
	}

	fmt.Println(order.String())
	fmt.Println(time.Since(start))
}

func (s sortNodes) Len() int {
	return len(s)
}

func (s sortNodes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortNodes) Less(i, j int) bool {
	return s[i].id < s[j].id
}
