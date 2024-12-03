package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var DiffCount = 0

func main() {
	file, err := os.Open("test_input.txt")
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

	slices.Sort(list_1)
	slices.Sort(list_2)

	for i := 0; i < len(list_2); i++ {
		DiffCount += absDiffInt(list_1[i], list_2[i])
	}
	
	fmt.Println(DiffCount)
}

func absDiffInt(x, y int) int {
	if x < y {
	   return y - x
	}
	return x - y
 }