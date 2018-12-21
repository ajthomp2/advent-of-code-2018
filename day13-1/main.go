package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type direction int

const (
	UP direction = 1
	DOWN direction = 2
	LEFT direction = 3
	RIGHT direction = 4
	STRAIGHT direction = 5
)

type cart struct {
	id int
	row, column int
	dir direction
	replaceChar byte
	turnDir direction
}

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal("error reading data")
	}

	// create track and carts from input
	scanner := bufio.NewScanner(file)
	var track []string
	var carts []cart
	var id int
	for scanner.Scan() {
		line := scanner.Text()
		track = append(track, line)
		for i, char := range line {
			switch char {
			case 'v':
				c := cart{
					id: id,
					row: len(track) - 1,
					column: i,
					dir: DOWN,
					replaceChar: '|',
					turnDir: LEFT,
				}
				id++
				carts = append(carts, c)
			case '>':
				c := cart{
					id: id,
					row: len(track) - 1,
					column: i,
					dir: RIGHT,
					replaceChar: '-',
					turnDir: LEFT,
				}
				id++
				carts = append(carts, c)
			case '<':
				c := cart{
					id: id,
					row: len(track) - 1,
					column: i,
					dir: LEFT,
					replaceChar: '-',
					turnDir: LEFT,
				}
				id++
				carts = append(carts, c)
			case '^':
				c := cart{
					id: id,
					row: len(track) - 1,
					column: i,
					dir: UP,
					replaceChar: '|',
					turnDir: LEFT,
				}
				id++
				carts = append(carts, c)
			}
		}
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(carts)
	for i, row := range track {
		fmt.Printf("%3d %s\n", i, row)
	}
	fmt.Println("----------------------------")

	time.Sleep(10 * time.Second)

	// move carts
	var count, crashx, crashy int
	for {
		var crash bool
		for i, c := range carts {
			var newChar string
			track[c.row] = track[c.row][:c.column] + string(c.replaceChar) + track[c.row][c.column+1:]
			switch c.dir {
			case UP:
				carts[i].row -= 1
				switch track[c.row-1][c.column] {
				case '|':
					newChar = "^"
					carts[i].replaceChar = '|'
				case '/':
					newChar = ">"
					carts[i].replaceChar = '/'
					carts[i].dir = RIGHT
				case '\\':
					newChar = "<"
					carts[i].replaceChar = '\\'
					carts[i].dir = LEFT
				case '+':
					carts[i].replaceChar = '+'
					if c.turnDir == LEFT {
						newChar = "<"
						carts[i].turnDir = STRAIGHT
						carts[i].dir = LEFT
					} else if c.turnDir == STRAIGHT {
						newChar = "^"
						carts[i].turnDir = RIGHT
					} else {
						newChar = ">"
						carts[i].turnDir = LEFT
						carts[i].dir = RIGHT
					}
				case '<', '^', '>', 'v':
					crash = true
					crashx = c.column
					crashy = c.row - 1
				}
				track[c.row-1] = track[c.row-1][:c.column] + newChar + track[c.row-1][c.column+1:]
			case DOWN:
				carts[i].row += 1

				switch track[c.row+1][c.column] {
				case '|':
					newChar = "v"
					carts[i].replaceChar = '|'
				case '/':
					newChar = "<"
					carts[i].replaceChar = '/'
					carts[i].dir = LEFT
				case '\\':
					newChar = ">"
					carts[i].replaceChar = '\\'
					carts[i].dir = RIGHT
				case '+':
					carts[i].replaceChar = '+'
					if c.turnDir == LEFT {
						newChar = ">"
						carts[i].turnDir = STRAIGHT
						carts[i].dir = RIGHT
					} else if c.turnDir == STRAIGHT {
						newChar = "v"
						carts[i].turnDir = RIGHT
					} else {
						newChar = "<"
						carts[i].turnDir = LEFT
						carts[i].dir = LEFT
					}
				case '<', '^', '>', 'v':
					crash = true
					crashx = c.column
					crashy = c.row + 1
				}
				track[c.row+1] = track[c.row+1][:c.column] + newChar + track[c.row+1][c.column+1:]
			case LEFT:
				carts[i].column -= 1

				switch track[c.row][c.column-1] {
				case '-':
					newChar = "<"
					carts[i].replaceChar = '-'
				case '/':
					newChar = "v"
					carts[i].replaceChar = '/'
					carts[i].dir = DOWN
				case '\\':
					newChar = "^"
					carts[i].replaceChar = '\\'
					carts[i].dir = UP
				case '+':
					carts[i].replaceChar = '+'
					if c.turnDir == LEFT {
						newChar = "v"
						carts[i].turnDir = STRAIGHT
						carts[i].dir = DOWN
					} else if c.turnDir == STRAIGHT {
						newChar = "<"
						carts[i].turnDir = RIGHT
					} else {
						newChar = "^"
						carts[i].turnDir = LEFT
						carts[i].dir = UP
					}
				case '<', '^', '>', 'v':
					crash = true
					crashx = c.column - 1
					crashy = c.row
				}
				track[c.row] = track[c.row][:c.column-1] + newChar + track[c.row][c.column:]
			case RIGHT:
				carts[i].column += 1

				switch track[c.row][c.column+1] {
				case '-':
					carts[i].replaceChar = '-'
					newChar = ">"
				case '/':
					newChar = "^"
					carts[i].replaceChar = '/'
					carts[i].dir = UP
				case '\\':
					newChar = "v"
					carts[i].replaceChar = '\\'
					carts[i].dir = DOWN
				case '+':
					carts[i].replaceChar = '+'
					if c.turnDir == LEFT {
						newChar = "^"
						carts[i].turnDir = STRAIGHT
						carts[i].dir = UP
					} else if c.turnDir == STRAIGHT {
						newChar = ">"
						carts[i].turnDir = RIGHT
					} else {
						newChar = "v"
						carts[i].turnDir = LEFT
						carts[i].dir = DOWN
					}
				case '<', '^', '>', 'v':
					crash = true
					crashx = c.column + 1
					crashy = c.row
				}
				track[c.row] = track[c.row][:c.column+1] + newChar + track[c.row][c.column+2:]
			}
		}

		fmt.Println(carts)
		//for _, c1 := range carts {
		//	for _, c2 := range carts {
		//		if c1.id == c2.id {
		//			continue
		//		}
		//		if c1.row == c2.row && c1.column == c2.column {
		//			fmt.Println(c1, c2)
		//			crash = true
		//		}
		//	}
		//}
		for i, row := range track {
			fmt.Printf("%3d %s\n", i, row)
		}
		fmt.Println(count)
		//fmt.Println()
		fmt.Println("-------------------------------------------------------------------------------------")

		count++


		if crash {
			break
		}
		time.Sleep(time.Second)
	}

	fmt.Printf("%v,%v", crashx, crashy)
}