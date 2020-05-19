/*
Binary Search Algorithm in Go
Author: John M. Hamm
*/

package main

import "fmt"

func Binarysearch(data []int, value int) bool {
	size := len(data)
	low := 0
	high := size - 1
	mid := 0 //we dynamically set this

	for low <= high {
		mid = (low + high) / 2
		if data[mid] == value {
			return true
		} else if data[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

func main() {
	arr := []int{2, 4, 6, 8, 10, 12, 14, 18, 21, 23, 24}
	fmt.Println(Binarysearch(arr, 18)) //true its there
	fmt.Println(Binarysearch(arr, 11)) //false its there
}
