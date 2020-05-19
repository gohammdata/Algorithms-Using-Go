/*
Insertion Sort Algorithm in Go
File: insertionSort.go
Author: John M. Hamm
Purpose: Stanford Algorithms Course.
*/

//A bit quicker than Bubble Sort but still at O(n^2) complexity
//Common anectdote is this is how most people sort playing cards in a hand.

package main //main for now to run each. will change to file name on build and fun from main.go

import "fmt"

func InsertionSort(arr []int, comp func(int, int) bool) {
	size := len(arr)
	var temp, i, j int
	for i = 1; i < size; i++ {
		temp = arr[i]
		for j = i; j > 0 && comp(arr[j-1], temp); j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = temp
	}
}

func more(value1 int, value2 int) bool {
	return value1 > value2
}

func main() {
	data := []int{7, 9, 2, 5, 8, 1, 3, 4, 6}
	InsertionSort(data, more)
	fmt.Println(data)
}
