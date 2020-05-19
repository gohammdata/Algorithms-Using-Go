package main

import "fmt"

func linearSearchUnsorted(data []int, value int) bool {
	size := len(data)
	for i := 0; i < size; i++ {
		if value == data[i] {
			return true
		}
	}
	return false
}

func main() {
	arr := []int{2, 4, 6, 8, 10, 12, 14, 18, 21, 23, 24}
	fmt.Println(linearSearchUnsorted(arr, 18)) //true
	fmt.Println(linearSearchUnsorted(arr, 11)) //false not there
}
