package main

import (
	"fmt"
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	stack := []int{}
	res := make([]int, n)
	prevTime := 0

	for _, log := range logs {
		parsed := strings.Split(log, ":")
		id, _ := strconv.Atoi(parsed[0])
		time, _ := strconv.Atoi(parsed[2])
		event := parsed[1]

		difft := time - prevTime

		switch event {
		case "start":
			if len(stack) > 0 {
				res[stack[len(stack)-1]] += difft
			}
			stack = append(stack, id)
			prevTime = time
		case "end":
			res[stack[len(stack)-1]] += difft + 1

			stack = stack[:len(stack)-1]

			prevTime = time + 1
		}
	}

	return res
}

func main() {
	n := 2

	logs := []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"}

	fmt.Println(exclusiveTime(n, logs))

}
