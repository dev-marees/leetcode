package main

import "fmt"

func main() {
	nums := []int{8, 1, 2, 2, 3}
	nums1 := []int{7, 7, 7, 7}
	nums2 := []int{6, 5, 4, 8}
	fmt.Println("counts of smaller", smallerNumbersThanCurrent(nums))
	fmt.Println("counts of smaller", smallerNumbersThanCurrent(nums1))
	fmt.Println("counts of smaller", smallerNumbersThanCurrent(nums2))
}

func smallerNumbersThanCurrent(nums []int) []int {
	result := []int{}
	for i := 0; i < len(nums); i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			if nums[i] > nums[j] {
				count++
			}
		}
		result = append(result, count)
	}
	return result
}
