package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const (
	UP    = "UP"
	DOWN  = "DOWN"
	LEFT  = "LEFT"
	RIGHT = "RIGHT"
)

var PatrolGrid = make(map[int][]string)

var ZeroIndexedGridHeight int
var ZeroIndexedGridWidth int 

func main() {
	
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to read file: %s", err)
	}

	y_index := 0
	for scanner.Scan() {
		currentLine := scanner.Text()
		for x_index, char := range currentLine {
			if string(char) == "^" {
				// fmt.Println("Guard details", Guard)
				Guard.setGuardDetails(y_index, x_index, UP)
				// fmt.Println("Guard details", Guard)
			}
			PatrolGrid[y_index] = append(PatrolGrid[y_index], string(char))
		}
		y_index++
	}
	
	
	ZeroIndexedGridHeight = len(PatrolGrid) -1
	ZeroIndexedGridWidth = len(PatrolGrid[0]) -1

	for {
		// get next possible block
		next_y_index, next_x_index := getNextPossiblePosition(Guard.direction, Guard.y_index, Guard.x_index)
		if isOutOfBounds(next_y_index, next_x_index) {
			// fmt.Println("out of bounds exiting")
			break
		}
		if isObstacle(next_y_index, next_x_index) {
			// fmt.Println("obstacle at:", "y_index", next_y_index, "x_index", next_x_index)
			Guard.getRotateDirection90DegreesRight()	
			continue // now that we are going in a diffrent direction we want to check for obstacles and out of bounds again
		}	
		
		// move to next block
		Guard.advanceOneBlock(next_y_index, next_x_index)
		
		// no obstacle and not out of bounds move to next block
		// fmt.Println("after move")
		printGrid(PatrolGrid)
	}
	
	// fmt.Println("Cells visited by Guard:", len(Guard.visited))
	
	
	printGridsInOrder()
}

func (g *GuardDetails) advanceOneBlock(next_y_index int, next_x_index int) {
	PatrolGrid[g.y_index][g.x_index] = "X"
	g.y_index = next_y_index
	g.x_index = next_x_index
	g.visited[fmt.Sprintf("%d,%d", g.y_index, g.x_index)] = true
	PatrolGrid[Guard.y_index][Guard.x_index] = g.dirSymbol
}

func isOutOfBounds(y_index int, x_index int) bool {
	return y_index < 0 || y_index > ZeroIndexedGridHeight || x_index < 0 || x_index > ZeroIndexedGridWidth
}

func isObstacle(y_index int, x_index int) bool {
	return PatrolGrid[y_index][x_index] == "#"
}

func getNextPossiblePosition(direction string, y_index int, x_index int) (int, int) {

	switch direction {
	case UP:
		y_index-- 
	case DOWN:
		y_index++
	case LEFT:
		x_index--
	case RIGHT:
		x_index++
	}
	// fmt.Println("direction",direction,"y_index",y_index , "x_index", x_index) 
	return y_index,x_index
}

type GuardDetails struct {
	y_index   int
	x_index   int
	direction string
	dirSymbol string
	visited   map[string]bool
}

var Guard = GuardDetails{}

func (g *GuardDetails) setGuardDetails(y_index int, x_index int, direction string) {
	g.y_index = y_index
	g.x_index = x_index
	g.direction = direction
	g.dirSymbol = "^"
	g.visited = make(map[string]bool)
	g.visited[fmt.Sprintf("%d,%d", g.y_index, g.x_index)] = true // record the starting position
}

func (g *GuardDetails) getRotateDirection90DegreesRight() {
	// fmt.Println("getRotateDirection90DegreesRight", "starting direction:", g.direction)
	switch g.direction {
	case UP:
		g.direction = RIGHT
		g.dirSymbol = ">"
	case RIGHT:	
		g.direction = DOWN
		g.dirSymbol = "v"
	case DOWN:		
		g.direction = LEFT
		g.dirSymbol = "<"
	case LEFT:		
		g.direction = UP
		g.dirSymbol = "^"
	}
	// fmt.Println("getRotateDirection90DegreesRight", "end direction:", g.direction)
}

var GridStringsInOrder = []string{}
var VisitedCellsCount = []int{}

func printGrid(grid map[int][]string) {
	
	gridString := ""
	for i := 0; i < len(grid); i++ {
		gridString += strings.Join(grid[i], "") + "\n" 
	}
	
	GridStringsInOrder = append(GridStringsInOrder, gridString)
	VisitedCellsCount = append(VisitedCellsCount, len(Guard.visited))
}

func printGridsInOrder() {
	ticker := time.Tick(time.Millisecond * 1)
	for i, gridString := range GridStringsInOrder {
		<-ticker
		fmt.Printf("\033[0;0H")
		fmt.Printf("\x0cVisit Count: %d\n", VisitedCellsCount[i])
		fmt.Printf("\x0c %s", gridString)
	}
}