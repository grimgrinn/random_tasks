package main

import "fmt"

// sam
func lastStoneWeight1(stones []int) int {
	fmt.Println(stones)
	for j := len(stones); j > 0; j-- {
		heavy1 := 0
		heavy2 := 0

		heavyIndex1 := 0
		heavyIndex2 := 0

		if len(stones) == 0 {
			return 0
		}
		if len(stones) == 1 {
			return stones[0]
		}

		for i := 0; i < len(stones); i++ {
			if stones[i] > heavy1 {
				heavy1 = stones[i]
				heavyIndex1 = i
			}
		}

		for i := 0; i < len(stones); i++ {
			if stones[i] > heavy2 && i != heavyIndex1 {
				heavy2 = stones[i]
				heavyIndex2 = i
			}
		}

		if heavy1 == heavy2 {
			if heavyIndex1 > heavyIndex2 {
				stones = append(stones[:heavyIndex1], stones[heavyIndex1+1:]...)
				stones = append(stones[:heavyIndex2], stones[heavyIndex2+1:]...)
			} else {
				stones = append(stones[:heavyIndex2], stones[heavyIndex2+1:]...)
				stones = append(stones[:heavyIndex1], stones[heavyIndex1+1:]...)
			}
		} else {
			stones[heavyIndex1] = heavy1 - heavy2
			stones = append(stones[:heavyIndex2], stones[heavyIndex2+1:]...)
		}

		fmt.Println(stones)

		fmt.Println(heavyIndex1, heavyIndex2)
		fmt.Println(heavy1, heavy2)

	}

	return 0
}

// // sprosil
// func lastStoneWeight(stones []int) int {
// 	for len(stones) > 1 {
// 		idx1, idx2 := 0, -1
// 		for i := 1; i < len(stones); i++ {
// 			if stones[i] > stones[idx1] {
// 				idx2 = idx1
// 				idx1 = i
// 			} else if idx2 == -1 || stones[i] > stones[idx2] {
// 				idx2 = i
// 			}
// 		}

// 		diff := stones[idx1] - stones[idx2]
// 		if diff < 0 {
// 			diff = -diff
// 		}

// 		if idx1 > idx2 {
// 			stones[idx2] = diff
// 			stones[idx1] = stones[len(stones)-1]
// 		} else {
// 			stones[idx1] = diff
// 			stones[idx2] = stones[len(stones)-1]
// 		}
// 		stones = stones[:len(stones)-1]
// 	}

// 	if len(stones) == 1 {
// 		return stones[0]
// 	}

// 	return 0
// }

// heap
func siftUp(heap []int, i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if heap[i] > heap[parent] {
			heap[i], heap[parent] = heap[parent], heap[i]
			i = parent
		} else {
			break
		}
	}
}

func siftDown(heap []int, i int, n int) {
	for {
		left := 2*i + 1
		right := 2*i + 2
		largest := i

		if left < n && heap[left] > heap[largest] {
			largest = left
		}
		if right < n && heap[right] > heap[largest] {
			largest = right
		}

		if largest == i {
			break
		}

		heap[i], heap[largest] = heap[largest], heap[i]

		i = largest
	}
}

func lastStoneWeight(stones []int) int {
	for i := len(stones)/2 - 1; i >= 0; i-- {
		siftDown(stones, i, len(stones))
	}

	for len(stones) > 1 {
		heavy1 := stones[0]

		stones[0] = stones[len(stones)-1]
		stones = stones[:len(stones)-1]

		siftDown(stones, 0, len(stones))

		heavy2 := stones[0]

		stones[0] = stones[len(stones)-1]
		stones = stones[:len(stones)-1]
		siftDown(stones, 0, len(stones))

		if heavy1 != heavy2 {
			stones = append(stones, heavy1-heavy2)
			siftUp(stones, len(stones)-1)
		}
	}

	if len(stones) == 1 {
		return stones[0]
	}
	return 0
}

func main() {
	stones := []int{2, 7, 4, 1, 8, 1}

	fmt.Println(lastStoneWeight(stones))

}
