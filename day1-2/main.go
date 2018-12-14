package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	filename := "data.txt"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("error reading file")
	}

	values := strings.Split(string(data), "\r\n")

	var nums []int
	for _, v := range values {
		num, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("error converting string %v: %v", v, err)
		}
		nums = append(nums, num)
	}

	freqs := make(map[int]struct{})
	var sum int
	done := false
	for {
		for _, num := range nums {
			sum += num
			if _, ok := freqs[sum]; ok {
				done = true
				break
			}
			freqs[sum] = struct{}{}
		}
		if done {
			break
		}
	}
	fmt.Println(sum)
}
