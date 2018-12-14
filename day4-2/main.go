package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

const layout = "2006-01-02 15:04"

type byTimeStamp []string

func main() {
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal("unable to read data from file")
	}

	inputs := strings.Split(string(data), "\r\n")
	sort.Sort(byTimeStamp(inputs))

	scheds := make(map[int][60]int)

	var currId int
	var fellAsleep time.Time
	for _, input := range inputs {
		if strings.Contains(input, "Guard #") {
			currIdStr := strings.Split(strings.Split(input, "#")[1], " ")[0]
			currId, err = strconv.Atoi(currIdStr)
			if err != nil {
				log.Fatal("error parsing ID string: ", currIdStr)
			}
		} else if strings.Contains(input, "falls asleep") {
			fellAsleep = parseShiftTime(input)
		} else {
			t := parseShiftTime(input)
			sched := scheds[currId]
			for i := fellAsleep.Minute(); i < t.Minute(); i++ {
				sched[i]++
			}
			scheds[currId] = sched
		}
	}

	var id, maxMin, maxMinInd int
	for k, v := range scheds {
		for i, a := range v {
			if a > maxMin {
				maxMin = a
				maxMinInd = i
				id = k
			}
		}
	}

	fmt.Println(maxMinInd * id)
}

// sorting interface functions
func (b byTimeStamp) Len() int {
	return len(b)
}

func (b byTimeStamp) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b byTimeStamp) Less(i, j int) bool {
	ti := parseShiftTime(b[i])
	tj := parseShiftTime(b[j])

	// check if diff is negative
	if diff := ti.Sub(tj); diff.String()[0] == '-' {
		return true
	}
	return false
}

func parseShiftTime(input string) time.Time {
	shiftTimeStr := strings.Split(strings.Split(string(input), "[")[1], "]")[0]
	shiftTime, err := time.Parse(layout, shiftTimeStr)
	if err != nil {
		log.Fatal("error parsing time: ", shiftTimeStr)
	}
	return shiftTime
}
