package main

import "fmt"

func getConcatenation(nums []int) []int {
	result := make([]int, len(nums))
	copy(result, nums)
	for i := 0; i < len(nums); i++ {
		result = append(result, nums[i])
	}

	return result

}

func main() {
	nums := []int{1, 2, 1}
	fmt.Println("Array concadination ::", getConcatenation(nums))
}
