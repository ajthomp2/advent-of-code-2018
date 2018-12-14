package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	filename := "data.txt"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("unable to read file")
	}

	ids := strings.Split(string(data), "\n")

	lcMap := make(map[int]int)
	for _, id := range ids {
		charMap := make(map[rune]int)
		for _, char := range id {
			charMap[char]++
		}
		addedMap := make(map[int]struct{})
		for _, v := range charMap {
			if v > 1 {
				if _, ok := addedMap[v]; !ok {
					lcMap[v]++
					addedMap[v] = struct{}{}
				}
			}
		}
	}
	checkSum := 1
	for _, v := range lcMap {
		checkSum *= v
	}
	fmt.Println(checkSum)
}
