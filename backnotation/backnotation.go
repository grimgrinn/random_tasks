package main

import (
	"fmt"
	"strconv"
)

func evalrpn(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		fmt.Println(token)
		digit, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, digit)
			continue
		}

		b := stack[len(stack)-1]
		a := stack[len(stack)-2]

		stack = stack[:len(stack)-2]

		var res int
		switch token {
		case "+":
			res = a + b
		case "-":
			res = a - b
		case "*":
			res = a * b
		case "/":
			res = a / b
		}
		stack = append(stack, res)

	}

	return stack[0]
}

func main() {
	fmt.Println("jopa")

	tokens := []string{"2", "1", "+", "3", "*"}

	result := evalrpn(tokens)

	fmt.Println("result:", result)
}
