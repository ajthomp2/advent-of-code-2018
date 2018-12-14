package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"time"
)

func main() {
	start := time.Now()
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := react(data)
	fmt.Println(len(result))
	fmt.Println(time.Since(start))
}

func react(polymer []byte) []byte {
	l := len(polymer)
	if l == 1 {
		return polymer
	}
	mid := l / 2
	return combine(react(polymer[:mid]), react(polymer[mid:]))
}

func combine(left, right []byte) []byte {
	if len(left) == 0 || len(right) == 0 {
		return append(left, right...)
	}
	if leftLastInd := len(left) - 1; math.Abs(float64(left[leftLastInd]) - float64(right[0])) == 32 {
		return combine(left[:leftLastInd], right[1:])
	}
	return append(left, right...)
}
