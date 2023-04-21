package main

import (
	"fmt"
)

// Bubble Sort
func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// Selection Sort
func selectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

// Insertion Sort
func insertionSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}

// Quick Sort
func quickSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	pivot := arr[n-1]
	left, right := []int{}, []int{}
	for i := 0; i < n-1; i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	left, right = quickSort(left), quickSort(right)
	return append(append(left, pivot), right...)
}

// Merge Sort
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left, right := mergeSort(arr[:mid]), mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	res := []int{}
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			res, left = append(res, left[0]), left[1:]
		} else {
			res, right = append(res, right[0]), right[1:]
		}
	}
	res = append(res, left...)
	res = append(res, right...)
	return res
}

func main() {
	arr := []int{64, 25, 12, 22, 11, 12, 9}

	// Bubble Sort
	fmt.Println("Bubble Sort:")
	fmt.Println(bubbleSort(arr))

	// Selection Sort
	fmt.Println("Selection Sort:")
	fmt.Println(selectionSort(arr))

	// Insertion Sort
	fmt.Println("Insertion Sort:")
	fmt.Println(insertionSort(arr))

	// Quick Sort
	fmt.Println("Quick Sort:")
	fmt.Println(quickSort(arr))

	// Merge Sort
	fmt.Println("Merge Sort:")
	fmt.Println(mergeSort(arr))
}
