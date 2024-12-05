package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var WordSearchLines = make(map[int][]string)

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
		for _, char := range currentLine {
			WordSearchLines[y_index] = append(WordSearchLines[y_index], string(char))
		}
		y_index++
	}

	fmt.Println(WordSearchLines)
	totalScore := 0
	totalScore += processDiagonally(WordSearchLines)
	totalScore += processHorizontally(WordSearchLines)
	totalScore += processVertically(WordSearchLines)

	fmt.Println("total score", totalScore)
}

func processHorizontally(WordSearchLines map[int][]string) int {
	scoreTotal := 0

	fmt.Println("horizontal scores")
	for _, line := range WordSearchLines {
		fmt.Println(strings.Join(line, ""), xmasCount(strings.Join(line, "")))
		scoreTotal += xmasCount(strings.Join(line, ""))
	}
	return scoreTotal
}

func processVertically(WordSearchLines map[int][]string) int {
	scoreTotal := 0

	VerticalLines := make(map[int]string)

	for y := 0; y < len(WordSearchLines); y++ {
		for i, char := range WordSearchLines[y] {
			VerticalLines[i] += string(char)
		}
	}
	fmt.Println("vertical scores")
	for _, line := range VerticalLines {
		fmt.Println(line, xmasCount(line))
		scoreTotal += xmasCount(line)
	}
	return scoreTotal
}

func processDiagonally(WordSearchLines map[int][]string) int {

	scoreTotal := 0

	DiagonalLines := parseDiagonalLines(WordSearchLines)

	fmt.Println("diagonal scores")
	for _, line := range DiagonalLines {
		fmt.Println(line, xmasCountNoFlip(line))
		scoreTotal += xmasCountNoFlip(line)
	}
	return scoreTotal
}

func parseDiagonalLines(diagonalLines map[int][]string) []string {
	var result []string

	// Get the size of the matrix
	size := len(diagonalLines)

	// Get all possible diagonal lines in all 4 directions
	for k := -size + 1; k < size; k++ {
		var diagonal1, diagonal2, diagonal3, diagonal4 string
		for i := 0; i < size; i++ {
			j := i + k
			if j >= 0 && j < size {
				diagonal1 += diagonalLines[i][j]
				diagonal2 += diagonalLines[i][size-1-j]
				diagonal3 += diagonalLines[size-1-i][j]
				diagonal4 += diagonalLines[size-1-i][size-1-j]
			}
		}
		result = append(result, diagonal1)
		result = append(result, diagonal2)
		result = append(result, diagonal3)
		result = append(result, diagonal4)
	}

	return result
}

func xmasCount(line string) int {
	xmasParts := strings.Split(line, "XMAS")
	samxParts := strings.Split(line, "SAMX")
	return (len(xmasParts) - 1) + (len(samxParts) - 1)
}


func xmasCountNoFlip(line string) int {
	xmasParts := strings.Split(line, "XMAS")
	return (len(xmasParts) - 1) 
}
