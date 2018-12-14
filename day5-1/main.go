package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	start := time.Now()
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal("unable to read data from file")
	}

	polymer := data
	for {
		l := len(polymer)
		buf, bufInd := make([]byte, l), 0
		done := true
		for i := 0; i < l; i++ {
			if i < l - 1 && isPair(polymer[i], polymer[i + 1]) {
				i++
				done = false
			} else {
				buf[bufInd] = polymer[i]
				bufInd++
			}
		}
		if done {
			buf = nil
			break
		}
		polymer = buf[:bufInd]
	}

	fmt.Println(len(polymer))
	fmt.Println(time.Since(start))
}

func isPair(left, right byte) bool {
	return left + 32 == right || left == right + 32
}