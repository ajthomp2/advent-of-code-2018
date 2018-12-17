package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
	"time"
)

const TASK_BASE_TIME = 60
const NUM_WORKERS = 5

type doubleLinkedNode struct {
	id string
	next []*doubleLinkedNode
	prev []*doubleLinkedNode
}

type sortNodes []*doubleLinkedNode

type worker struct {
	task *doubleLinkedNode
	timeLeft int
}

func main() {
	start := time.Now()

	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal("error reading data")
	}

	// create doubly linked node graph
	nodes := make(map[string]*doubleLinkedNode)
	for _, instruction := range strings.Split(string(data), "\r\n") {
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

	workers := [NUM_WORKERS]worker{}
	var total int
	for {
		sort.Sort(sortNodes(nextNodes))

		// assign tasks to workers
		var c int
		for _, n := range nextNodes {
			for i := range workers {
				if workers[i].task == nil {
					workers[i].task = n
					workers[i].timeLeft = TASK_BASE_TIME + (int(n.id[0]) - 64)
					c++
					break
				}
			}
		}
		nextNodes = nextNodes[c:]

		// get first task done
		var nextWorkerDone *worker
		for i := range workers {
			if workers[i].task != nil && (nextWorkerDone == nil || workers[i].timeLeft < nextWorkerDone.timeLeft) {
				nextWorkerDone = &workers[i]
			}
		}

		// end condition when no more workers working
		if nextWorkerDone == nil {
			break
		}

		for ci := range nextWorkerDone.task.next {
			for pi := range nextWorkerDone.task.next[ci].prev {
				if nextWorkerDone.task.next[ci].prev[pi].id == nextWorkerDone.task.id {
					nextWorkerDone.task.next[ci].prev = append(nextWorkerDone.task.next[ci].prev[:pi], nextWorkerDone.task.next[ci].prev[pi+1:]...)
					break
				}
			}
			if len(nextWorkerDone.task.next[ci].prev) == 0 {
				nextNodes = append(nextNodes, nextWorkerDone.task.next[ci])
			}
		}

		timeElapsed := nextWorkerDone.timeLeft
		for i := range workers {
			if workers[i].task != nil {
				workers[i].timeLeft -= timeElapsed
			}
		}

		total += timeElapsed
		nextWorkerDone.task = nil
	}

	fmt.Println(total)
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
