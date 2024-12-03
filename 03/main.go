package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const ClosingParen = 41
const Comma = 44

func main() {

	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	sParts := strings.Split(string(data), "mul(")

	MultiResult := 0

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

	fmt.Println("MultiResult: ", MultiResult)
}
