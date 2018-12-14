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

	inputs := strings.Split(string(data), "\r\n")

	var sum int
	for _, input := range inputs {
		num, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal("error converting string: ", input)
		}
		sum += num
	}
	fmt.Println(sum)
}
