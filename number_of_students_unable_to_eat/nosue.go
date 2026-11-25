package main

import "fmt"

// sam
func countStudents1(students []int, sandwiches []int) int {
	fmt.Println(students)
	c := 0
	for i := 0; len(sandwiches) <= len(students); i++ {
		if len(students) == 0 {
			return 0
		}

		s := students[0]
		fmt.Printf("i = %d; s = %d, len(students) = %d\n", i, s, len(students))
		students = students[1:]
		if s == sandwiches[0] {
			sandwiches = sandwiches[1:]
		} else {
			students = append(students, s)
			c++
			if c > 300 {
				break
			}
		}
		fmt.Println(students)
		fmt.Println(sandwiches)

	}

	return len(students)
}

// sprosil
func countStudents(students []int, sandwiches []int) int {
	want0, want1 := 0, 0
	for _, s := range students {
		switch s {
		case 0:
			want0++
		case 1:
			want1++
		}
	}

	for i := 0; i < len(sandwiches); i++ {
		if sandwiches[i] == 0 {
			if want0 == 0 {
				return want1
			}
			want0--
		} else {
			if want1 == 0 {
				return want0
			}
			want1--
		}
	}

	return 0
}

func main() {
	students := []int{1, 1, 1, 0, 0, 1}
	sandwiches := []int{1, 0, 0, 0, 1, 1}

	fmt.Println(countStudents(students, sandwiches))

}
