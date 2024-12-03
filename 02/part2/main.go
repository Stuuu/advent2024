package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

	var safeLvlCount = 0

	for scanner.Scan() {
		lvlStringParts := strings.Split(scanner.Text(), " ")

		shouldIncrease := intVal(lvlStringParts[0]) < intVal(lvlStringParts[1])
		safeLvl := true
		for i, lvl := range lvlStringParts {
			if i == 0 {
				continue
			}
			curLvlInt := intVal(lvl)
			lastLvlInt := intVal(lvlStringParts[i-1])
			if shouldIncrease {
				// 	The levels are either all increasing or all decreasing.
				// Any two adjacent levels differ by at least one and at most three.
				if !inRange(lastLvlInt, curLvlInt-3, curLvlInt-1) {
					safeLvl = false
					break
				}
			} else {
				if !inRange(lastLvlInt, curLvlInt+1, curLvlInt+3) {
					safeLvl = false
					break
				}
			}
		}
		if safeLvl {
			safeLvlCount++
		}
	}

	fmt.Println("safelvlCount:", safeLvlCount)
}

func intVal(char string) int {
	val, _ := strconv.Atoi(char)
	return val
}

func inRange(cur, min, max int) bool {
	return cur >= min && cur <= max
}
