package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const MAX_DIST = 10000

type zone struct {
	x, y     int
}

var zones []zone
var distMap [][]int
var maxX, maxY int
var area int

func main() {
	start := time.Now()
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal("error reading data")
	}

	coords := strings.Split(string(data), "\r\n")

	zones = make([]zone, len(coords))
	for i := range coords {
		zone := newZone(coords[i])
		if zone.x > maxX {
			maxX = zone.x
		}
		if zone.y > maxY {
			maxY = zone.y
		}
		zones[i] = zone
	}

	// create coordinate plane using max coordinate values
	distMap = make([][]int, maxX + 1)
	for i := range distMap {
		distMap[i] = make([]int, maxY + 1)
	}

	found := false
	for !found {
		rx := rand.Intn(maxX + 1)
		ry := rand.Intn(maxY + 1)
		if distMap[rx][ry] == 0 && totalDist(rx, ry) < MAX_DIST {
			visit(rx, ry)
			found = true
		}
	}

	fmt.Println(area)
	fmt.Println(time.Since(start))
}

func visit(x, y int) {
	if x < 0 || y < 0 || x > maxX || y > maxY {
		return
	}

	if distMap[x][y] != 0 {
		return
	}

	dist := totalDist(x, y)
	distMap[x][y] = dist
	if dist >= MAX_DIST {
		return
	}
	area++

	visit(x + 1, y)
	visit(x, y + 1)
	visit(x - 1, y)
	visit(x, y - 1)
}

func totalDist(x, y int) int {
	var total int
	for i := range zones {
		total += manhattanDist(zones[i].x, x, zones[i].y, y)
	}
	return total
}

func manhattanDist(x1, x2, y1, y2 int) int {
	return int(math.Abs(float64(x2 - x1)) + math.Abs(float64(y2 - y1)))
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
	}
	return zone
}
