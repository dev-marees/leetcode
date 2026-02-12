package main

import "fmt"

func main() {
	nums := []int{1, 0, 1, 1, 0, 1, 1, 1}
	fmt.Println("max consecutive 1's ::", findMaxConsecutiveOnes(nums))
}

func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	maxCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
			if maxCount < count {
				maxCount = count
			}
		} else {
			count = 0
		}
	}
	return maxCount
}
