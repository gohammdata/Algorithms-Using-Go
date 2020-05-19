/*
Simple Selection Sort Algorithm in Go.
File: selectionSort.go
Author: John M. Hamm
Purpose: Stanford Algorithms Course. See readme why Go.

Traveses an unsorted array to put the largest value at the end of it, repeated n-1 times.
Also quadratic complexity but better than bubble and insertion.
*/

package main //main for now to run each. will change to file name on build and fun from main.go

import "fmt"

func SelectionSort(arr []int) {
	size := len(arr)
	var i, j, max int
	for i = 0; i < size; i++ {
		max = 0
		for j = 1; j < size-1-i; j++ {
			if arr[j] > arr[max] {
				max = j
			}

		}
		arr[size-1-i], arr[max] = arr[max], arr[size-1-i]
	}
}

func main() {
	data := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	SelectionSort(data)
	fmt.Println(data)
}
