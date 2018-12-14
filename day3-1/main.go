package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type claim struct {
	id int
	offsetLeft, offsetTop int
	width, height int
}

func main() {
	filename := "data.txt"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("unable to read file")
	}

	claimsInput := strings.Split(string(data), "\r\n")
	claims := createClaims(claimsInput)

	sheet := createSheet(getMaxWidthAndHeight(claims))

	for _, c := range claims {
		for i := c.offsetTop; i < c.offsetTop + c.height; i++ {
			for j := c.offsetLeft; j < c.offsetLeft + c.width; j++ {
				sheet[j][i]++
			}
		}
	}

	var count int
	for _, row := range sheet {
		for _, val := range row {
			if val > 1 {
				count++
			}
		}
	}

	fmt.Println(count)
}

func createClaims(claimsInput []string) ([]claim) {
	var claims []claim

	for _, c := range claimsInput {
		parts := strings.Split(c, " ")

		// id
		id, err := strconv.Atoi(strings.Replace(parts[0], "#", "", 1))
		if err != nil {
			log.Fatalf("error converting %v to an int: %v", strings.Replace(parts[0], "#", "", 1), err)
		}

		// offsets
		offsets := strings.Split(strings.Replace(parts[2], ":", "", 1), ",")
		offsetLeft, err := strconv.Atoi(offsets[0])
		if err != nil {
			log.Fatalf("1 error converting %v to an int: %v", offsets[0])
		}
		offsetTop, err := strconv.Atoi(offsets[1])
		if err != nil {
			log.Fatalf("2 error converting %v to an int: %v", offsets[1])
		}

		// coordinates
		coords := strings.Split(parts[3], "x")
		width, err := strconv.Atoi(coords[0])
		if err != nil {
			log.Fatalf("3 error converting %v to an int: %v", coords[0])
		}
		height, err := strconv.Atoi(coords[1])
		if err != nil {
			log.Fatalf("3 error converting %v to an int: %v", coords[0])
		}

		// create claim
		newClaim := claim{
			id: id,
			offsetLeft: offsetLeft,
			offsetTop: offsetTop,
			width: width,
			height: height,
		}
		claims = append(claims, newClaim)
	}

	return claims
}

func getMaxWidthAndHeight(claims []claim) (int, int) {
	var maxWidth, maxHeight int
	for i := range claims {
		width := claims[i].offsetLeft + claims[i].width
		if width > maxWidth {
			maxWidth = width
		}

		height := claims[i].offsetTop + claims[i].height
		if height > maxHeight {
			maxHeight = height
		}
	}
	return maxWidth, maxHeight
}

func createSheet(width, height int) ([][]int) {
	sheet := make([][]int, width)
	for i := range sheet {
		sheet[i] = make([]int, height)
	}
	return sheet
}