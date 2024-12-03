package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main() {

	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	var doStringBuilder strings.Builder

	input_string_parts := string(data[:])
	
	NotInDontBlock := true
	for i, char := range input_string_parts {
		
		if i > 7 {
			if input_string_parts[i - 7 :i] == "don't()" {
				fmt.Println("don't hit", i)
				NotInDontBlock = false
			}
			if input_string_parts[i - 4 :i] == "do()" {
				fmt.Println("do hit", i)
				NotInDontBlock = true
			}
		}
		
		if NotInDontBlock {
			doStringBuilder.WriteString(string(char))
		}
		
	}
	fmt.Println("MultiResult: ",doMulParser(doStringBuilder.String()))
}

func doMulParser(data string) int {
	sParts := strings.Split(string(data), "mul(")

	MultiResult := 0
	
	const ClosingParen = 41
	const Comma = 44

	for _, part := range sParts {
		num1 := ""
		num2 := ""
		commaSeen := false
		for _, char := range part {

			if char == Comma {
				commaSeen = true
				continue
			}

			if char == ClosingParen {

				num1Int, _ := strconv.Atoi(num1)
				num2Int, _ := strconv.Atoi(num2)

				MultiResult += num1Int * num2Int
				break
			}
			if char < 47 || char > 58 {
				break
			}
			if !commaSeen {
				num1 += string(char)
			} else {
				num2 += string(char)
			}

		}
	}
	
	return	MultiResult
}