package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 4}
	nums2 := []int{1, 1}
	nums3 := []int{2, 2}
	fmt.Println("mis matched:", findErrorNums(nums))
	fmt.Println("mis matched:", findErrorNums(nums2))
	fmt.Println("mis matched:", findErrorNums(nums3))
}

func findErrorNums(nums []int) []int {

	temp := map[int]bool{}
	result := []int{}

	//actual sum
	actualSum := 0
	for i := 0; i < len(nums); i++ {
		actualSum += nums[i]
	}

	expectedSum := len(nums) * (len(nums) + 1) / 2
	for i := 0; i < len(nums); i++ {
		if _, ok := temp[nums[i]]; ok {
			missedNumber := expectedSum - (actualSum - nums[i])
			result = append(result, nums[i], missedNumber)
		}
		temp[nums[i]] = true
	}

	return result
}
