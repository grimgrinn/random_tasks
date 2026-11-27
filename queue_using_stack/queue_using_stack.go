package main

import "fmt"

// //sam

// type MyQueue struct {
// 	stack []int
// }

// func Counstructor() MyQueue {
// 	queue := new(MyQueue)
// 	return *queue
// }

// func (this *MyQueue) Push(x int) {
// 	this.stack = append(this.stack, x)
// }

// func (this *MyQueue) Pop() int {
// 	head := this.stack[0]
// 	this.stack = this.stack[1:]
// 	return head
// }

// func (this *MyQueue) Peek() int {
// 	head := this.stack[0]
// 	return head
// }

// func (this *MyQueue) Empty() bool {
// 	return len(this.stack) == 0
// }

type MyQueue struct {
	in  []int
	out []int
}

func Counstructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

func (this *MyQueue) shift() {
	if len(this.out) == 0 {
		for len(this.in) > 0 {
			top := this.in[len(this.in)-1]
			this.in = this.in[:len(this.in)-1]
			this.out = append(this.out, top)
		}
	}
}

func (this *MyQueue) Pop() int {
	this.shift()
	top := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	return top
}

func (this *MyQueue) Peek() int {
	this.shift()
	return this.out[len(this.out)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.in) == 0 && len(this.out) == 0
}

func main() {
	queue := Counstructor()

	fmt.Println(queue)

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)
	queue.Push(5)

	fmt.Println("queue - ", queue)

	fmt.Println("peek ", queue.Peek())

	fmt.Println("empty ", queue.Empty())

	fmt.Println("pop ", queue.Pop())

	fmt.Println("queue - ", queue)

	fmt.Println("pop ", queue.Pop())

	fmt.Println("queue - ", queue)

	fmt.Println("push 2")

	queue.Push(2)

	fmt.Println("queue - ", queue)

}
