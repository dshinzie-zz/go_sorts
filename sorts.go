package main

import (
	"fmt"
)

func insertionSort(items []int) []int {
	for i := 1; i < len(items); i++ {
		x := items[i] // current placeholder
		j := i        // position
		for j > 0 && items[j-1] > x {
			items[j] = items[j-1]
			j = j - 1
		}
		items[j] = x
	}
	return items
}

func bubbleSort(items []int) []int {
	swapped := true

	for swapped {
		swapped = false
		for i := 0; i < len(items)-1; i++ {
			if items[i] > items[i+1] {
				items[i], items[i+1] = items[i+1], items[i]
				swapped = true
			}
		}
	}
	return items
}

func mergeSort(items []int) []int {
	if len(items) <= 1 {
		return items
	}
	middle := (len(items)) / 2
	leftSlice := mergeSort(items[:middle])
	rightSlice := mergeSort(items[middle:])
	return merger(leftSlice, rightSlice)
}

func merger(left, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}

func main() {
	fmt.Println("---Starting sorts---")
	insertList := []int{5, 2, 9, 1, 13}
	bubbleList := []int{5, 2, 9, 1, 13}
	mergeList := []int{5, 2, 9, 1, 13}

	fmt.Println("---Insertion Sort---")
	fmt.Println(insertionSort(insertList))

	fmt.Println("---Bubble Sort---")
	fmt.Println(bubbleSort(bubbleList))

	fmt.Println("---Merge Sort---")
	fmt.Println(mergeSort(mergeList))

	fmt.Println("")
}
