package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	for i := range nums {
		fmt.Printf("Число - %d\n", nums[i])
	}
}
