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
	// store for all valid updates
	validUpdates := [][]string{}
	EmptyLineBreakForPageUpdates := false
	for scanner.Scan() {
		if scanner.Text() == "" {
			EmptyLineBreakForPageUpdates = true
			continue
		}

		if !EmptyLineBreakForPageUpdates {
			parseOrderRules(scanner.Text())
			continue
		}

		// parse update
		updateParts := strings.Split(scanner.Text(), ",")

		fmt.Println("BeforeRules", BeforeRules)
		fmt.Println("AfterRules", AfterRules)

		if isUpdateLineValid(updateParts) {
			validUpdates = append(validUpdates, updateParts)
			fmt.Println("updateParts is valid: ", updateParts)
		} else {
			fmt.Println("updateParts is invalid: ", updateParts)
		}
		
	}
	fmt.Println("Valid Updates", validUpdates)

	middleNumSum := 0
	for _, validUpdate := range validUpdates {
		fmt.Println("Valid Update: ", validUpdate)
		middleIndex := len(validUpdate) / 2
		fmt.Println("Middle Index: ", middleIndex)
		middleNum, _ := strconv.Atoi(validUpdate[middleIndex])
		middleNumSum += middleNum
	}
	
	fmt.Println("Middle Num Sum: ", middleNumSum)
}

func isUpdateLineValid(updateParts []string) bool {

	for updateIndex, updatePart := range updateParts {
		beforeParts := updateParts[:updateIndex]
		afterParts := updateParts[updateIndex+1:]

		fmt.Println("Before Parts", beforeParts)
		fmt.Println("After Parts", afterParts)
		// check all before rules for current updatePart

		fmt.Println("updatePart", updatePart)
		fmt.Println("BeforeRules", BeforeRules[updatePart])
		fmt.Println("AfterRules", AfterRules[updatePart])
		//1 check that each before part is not in not in after rules
		for _, beforeRule := range BeforeRules[updatePart] {
			if slices.Contains(beforeParts, beforeRule) {
				fmt.Println("updatePart does not meet before rule part: ", updatePart)
				fmt.Println("UpdatePart: ", updatePart, "Can't be before: ", beforeRule)
				fmt.Println("Based on BeforeRules: ", BeforeRules[updatePart])
				fmt.Println("Skipping this updateParts line", updateParts)
				return false
			}
		}

		//2 check that each after part is not in after rules
		for _, afterRule := range AfterRules[updatePart] {
			if slices.Contains(afterParts, afterRule) {
				fmt.Println("updatePart does not meet after rule part: ", updatePart)
				fmt.Println("UpdatePart: ", updatePart, "Can't be before: ", afterRule)
				fmt.Println("Based on AfterRules: ", AfterRules[updatePart])
				fmt.Println("Skipping this updateParts line", updateParts)

				return false
			}
		}
	}
	return true
}

var BeforeRules = make(map[string][]string)
var AfterRules = make(map[string][]string)

func parseOrderRules(rule string) {
	ruleParts := strings.Split(rule, "|")
	BeforeRules[ruleParts[0]] = append(BeforeRules[ruleParts[0]], ruleParts[1])
	AfterRules[ruleParts[1]] = append(AfterRules[ruleParts[1]], ruleParts[0])
}
