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
		log.Fatal("error reading file")
	}

	ids := strings.Split(string(data), "\n")

	for _, id := range ids {
		for _, otherId := range ids {
			var diff int
			for i := 0; i < len(id); i++ {
				if id[i] != otherId[i] {
					diff++
				}
			}
			if diff == 1 {
				fmt.Println(id, otherId)
			}
		}
	}
}