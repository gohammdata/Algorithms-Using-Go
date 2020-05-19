/*
Improved Bubble Sort Algorithm in Go.

Runtime is still the same for worst case, and average case O(n^2).
Nothing within the confines of a Bubble Sort can change these cases.
However, in this imporved Bubble Sort the runtime for Best case is now improved to O(n).

File: improvedbubbleSort.go
Author: John M. Hamm
Purpose: Stanford Algorithms Course. Improving bubbleSort.go to run quicker in a best case (same elsewhere).
*/

package main //main for now to run each. will change to file name on build and fun from main.go

import "fmt" //for print.

func ImporvedBubbleSort(arr []int, comp func(int, int) bool) {
	size := len(arr)
	swapped := 1

	for i := 0; i < (size-1) && swapped == 1; i++ {
		swapped = 0
		for j := 0; j < size-i-1; j++ {
			if comp(arr[j], arr[j+1]) {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				swapped = 1
			}
		}
	}
}
func more(value1 int, value2 int) bool {
	return value1 > value2
}

func main() {
	data := []int{9, 1, 8, 2, 7, 3, 6, 4, 5} //Can be any. This is just to test.
	ImporvedBubbleSort(data, more)
	fmt.Println(data)
}
