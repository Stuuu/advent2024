package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var SimmCount = 0

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	list_1 := []int{}
	list_2 := []int{}
	for scanner.Scan() {
		line_parts := strings.Split(scanner.Text(), "   ")
		i_1, _ := strconv.Atoi(line_parts[0])
		i_2, _ := strconv.Atoi(line_parts[1])
		list_1 = append(list_1, i_1)
		list_2 = append(list_2, i_2)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to read file: %s", err)
	}

	list2NumCounts := getNumCounts(list_2)

	for _, num := range list_1 {
		val, ok := list2NumCounts[num]
		if ok {
			SimmCount += num * val
		}
	}
	fmt.Println(SimmCount)
}

func getNumCounts(listOfNums []int) map[int]int {
	l2Counts := map[int]int{}
	for _, num := range listOfNums {
		_, ok := l2Counts[num]

		if !ok {
			l2Counts[num] = 1
		} else {
			l2Counts[num]++
		}
	}

	return l2Counts
}
