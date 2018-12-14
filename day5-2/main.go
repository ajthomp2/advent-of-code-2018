package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyz"

func main() {
	start := time.Now()
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal("unable to read data from file")
	}

	minLength := len(data)
	for _, letter := range letters {
		re := regexp.MustCompile("(?i)" + string(letter))
		polymer := []byte(re.ReplaceAllLiteralString(string(data), ""))

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
		if len(polymer) < minLength {
			minLength = len(polymer)
		}
	}

	fmt.Println(minLength)
	fmt.Println(time.Since(start))
}

func isPair(left, right byte) bool {
	return left + 32 == right || left == right + 32
}