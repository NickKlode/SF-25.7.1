package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	for i := range nums {
		fmt.Printf("Число - %d, Квадрат чис - %d\n", nums[i], nums[i]*nums[i])

	}
}
