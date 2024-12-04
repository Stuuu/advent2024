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

		unsafeLvls := getUnsafeLvls(lvlStringParts)

		if len(unsafeLvls) == 0 {
			safeLvlCount++
			// fmt.Println("safe lvl", lvlStringParts)
			continue
		}

		for i := 0; i <= len(lvlStringParts); i++ {
			unsafeRemoved := removeSliceIndex(lvlStringParts, i)
			if len(getUnsafeLvls(unsafeRemoved)) == 0 {
				safeLvlCount++
				// fmt.Println("safe lvl", lvlStringParts)
				// fmt.Println("unsafe corrected i: ", i, unsafeRemoved)
				break
			}
		}
		fmt.Println("unfixable", lvlStringParts)
	}

	fmt.Println("safelvlCount:", safeLvlCount)
}

func removeSliceIndex(slice []string, s int) []string {
	newSlice := []string{}
	for i, val := range slice {
		if i != s {
			newSlice = append(newSlice, val)
		}
	}
	return newSlice
}

func getUnsafeLvls(lvlStringParts []string) []int {

	badLvls := []int{}

	shouldIncrease := intVal(lvlStringParts[0]) < intVal(lvlStringParts[1])
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
				badLvls = append(badLvls, i)
			}
		} else {
			if !inRange(lastLvlInt, curLvlInt+1, curLvlInt+3) {
				badLvls = append(badLvls, i)
			}
		}
	}

	return badLvls
}

func intVal(char string) int {
	val, _ := strconv.Atoi(char)
	return val
}

func inRange(cur, min, max int) bool {
	return cur >= min && cur <= max
}
