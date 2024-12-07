package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	TestValue       int
	NumCombinations []int
	ComboCount      int
}

var Equations = []Equation{}
var UniqueCombos = map[int][]string{}

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

	uniqueComboCounts := map[int]bool{}
	for scanner.Scan() {
		equationParts := strings.Split(scanner.Text(), " ")

		intEquationParts := []int{}
		for _, equationPart := range equationParts {
			part := strings.TrimRight(equationPart, ":")
			intPart, _ := strconv.Atoi(part)
			intEquationParts = append(intEquationParts, intPart)
		}
		equation := Equation{}
		equation.Build(intEquationParts)
		Equations = append(Equations, equation)

		uniqueComboCounts[equation.ComboCount] = true
	}

	chars := []rune{'+', '*'}
	for comboCount, _ := range uniqueComboCounts {
		product := cartesianProduct(chars, comboCount)
		for _, combination := range product {
			UniqueCombos[comboCount] = append(UniqueCombos[comboCount], string(combination))
		}
	}

	totalCalibrationSum := 0
	for _, equation := range Equations {
		if equation.Test() {
			totalCalibrationSum += equation.TestValue
			fmt.Println(equation)
		}
	}
	fmt.Println(totalCalibrationSum)
}

func (equation *Equation) Test() bool {
	for _, combination := range UniqueCombos[equation.ComboCount] {
		runningTotal := equation.NumCombinations[0]
		combParts := strings.Split(combination, "")
		for i, num := range equation.NumCombinations[1:] {
			switch combParts[i] {
			case "+":
				runningTotal += num
			case "*":
				runningTotal *= num
			}
		}
		
		if runningTotal == equation.TestValue {
			fmt.Println("Equation Passed", equation.TestValue, combination)
			return true
		}
	}
	return false
}

func (equation *Equation) Build(equations []int) *Equation {
	equation.TestValue = equations[0]
	equation.NumCombinations = equations[1:]
	equation.ComboCount = len(equation.NumCombinations) - 1 // since operators are inbetween only

	return equation
}

func cartesianProduct(chars []rune, length int) [][]rune {
	if length == 1 {
		result := make([][]rune, len(chars))
		for i, char := range chars {
			result[i] = []rune{char}
		}
		return result
	}

	result := [][]rune{}
	for _, char := range chars {
		for _, product := range cartesianProduct(chars, length-1) {
			newProduct := make([]rune, len(product)+1)
			newProduct[0] = char
			copy(newProduct[1:], product)
			result = append(result, newProduct)
		}
	}
	return result
}
