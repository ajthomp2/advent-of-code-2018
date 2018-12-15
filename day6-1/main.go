package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

type location struct {
	id   int
	dist int
	seen bool
}

type zone struct {
	x, y     int
	count int
	infinite bool
}

var locationMap [][]location
var maxX, maxY int

func main() {
	start := time.Now()
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal("error reading data")
	}

	// create map of id to zone
	zones := make(map[int]*zone)
	for i, d := range strings.Split(string(data), "\r\n") {
		zone := newZone(d)
		if zone.x > maxX {
			maxX = zone.x
		}
		if zone.y > maxY {
			maxY = zone.y
		}
		zones[i+1] = &zone
	}

	// create coordinate plane using max coordinate values
	locationMap = make([][]location, maxY + 1)
	for i := range locationMap {
		locationMap[i] = make([]location, maxX + 1)
	}

	// populate coordinate plane with zone coordinates
	for k, v := range zones {
		locationMap[v.y][v.x] = location{id: k, dist: 0}
	}

	// find nearest zone for each point on coordinate plane
	for k, v := range zones {
		// reset all points to unseen
		for i := range locationMap {
			for j := range locationMap[i] {
				locationMap[i][j].seen = false
			}
		}
		visit(v.x, v.y, v.x, v.y, k)
	}

	// get counts for each zone
	for i := range locationMap {
		for j := range locationMap[i] {
			l := locationMap[i][j]
			if l.id != 0 {
				if i == 0 || j == 0 || i == maxY || j == maxX {
					zones[l.id].infinite = true
				}
				zones[l.id].count++
			}
		}
	}

	var max int
	for _, v := range zones {
		if v.infinite == false && v.count > max {
			max = v.count
		}
	}
	fmt.Println(max)
	fmt.Println(time.Since(start))
}

func visit(homex, homey, x, y, id int) {
	if x < 0 || y < 0 || x > maxX || y > maxY {
		return
	}

	l := locationMap[y][x]
	if l.seen {
		return
	}
	l.seen = true

	dist := int(math.Abs(float64(homex - x)) + math.Abs(float64(homey - y)))
	if dist < l.dist || (l.id == 0 && l.dist == 0) {
		l.id = id
		l.dist = dist
	} else if l.dist != 0 && dist == l.dist {
		// means no zone owns this
		l.id = 0
	} else if dist > l.dist {
		return
	}

	locationMap[y][x] = l

	visit(homex, homey, x + 1, y, id)
	visit(homex, homey, x, y + 1, id)
	visit(homex, homey, x - 1, y, id)
	visit(homex, homey, x, y - 1, id)
}

func newZone(coordsStr string) zone {
	coords := strings.Split(coordsStr, ", ")
	x, err := strconv.Atoi(coords[0])
	if err != nil {
		log.Fatal("error converting to int", coords[0])
	}
	y, err := strconv.Atoi(coords[1])
	if err != nil {
		log.Fatal("error converting to int", coords[1])
	}
	zone := zone{
		x:        x,
		y:        y,
		infinite: false,
	}
	return zone
}
