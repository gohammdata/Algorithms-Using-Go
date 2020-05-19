package main //main for now to run each. will change to file name on build and fun from main.go

import "fmt"

func BubbleSort(arr []int, comp func(int, int) bool) {
	size := len(arr)
	for i := 0; i < (size - 1); i++ {
		//Nested for loop resulting is one factor why bubble sorts have slow runtime for large data
		for j := 0; j < size-i-1; j++ {
			if comp(arr[j], arr[j+1]) {
				/* Swapping method which also is a bit slow as nested in 2 for loops */
				arr[j+1], arr[j] = arr[j], arr[j+1]

			}
		}
	}
}

func more(value1 int, value2 int) bool {
	return value1 > value2
}

func main() {
	data := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	BubbleSort(data, more)
	fmt.Println(data)
}
