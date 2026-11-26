package main

import "fmt"

func timeRequiredToBuy1(tickets []int, k int) int {
	//fmt.Println(tickets, k)

	count := 0
	i := 0
	for {
		fmt.Println(tickets)
		fmt.Printf("tickets[%d] - %d\n", i, tickets[i])

		head := tickets[0]
		tickets = tickets[1:]
		head -= 1

		count++

		if head > 0 {
			tickets = append(tickets, head)
		} else {
			if k > 0 {
				k -= 1
			}
		}

		if tickets[k] == 0 || k == 0 {
			break
		}

		if i < len(tickets)-1 {
			i++
		} else {
			i = 0
		}

	}

	return count
}

func timeRequiredToBuy(tickets []int, k int) int {
	count := 0
	for {
		tickets[0]--
		count++

		if k == 0 && tickets[0] == 0 {
			return count
		}

		if tickets[0] > 0 {
			tickets = append(tickets, tickets[0])
		}

		tickets = tickets[1:]

		if k == 0 {
			k = len(tickets) - 1
		} else {
			k--
		}
	}
}

func main() {
	tickets := []int{2, 3, 2}
	k := 2

	fmt.Println(timeRequiredToBuy(tickets, k))
}
