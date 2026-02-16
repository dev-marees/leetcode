package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	nums1 := []int{1, 1}
	fmt.Println("disappeared number", findDisappearedNumbers(nums))
	fmt.Println("disappeared number", findDisappearedNumbers(nums1))

}

func findDisappearedNumbers(nums []int) []int {
	result := []int{}
	//actual sum
	arrayLen := len(nums)
	temp := map[int]bool{}

	sort.Ints(nums)
	for i := 0; i < arrayLen; i++ {
		temp[nums[i]] = true
	}

	for i := 1; i <= arrayLen; i++ {
		if temp[i] {
			continue
		} else {
			result = append(result, i)
		}
	}

	return result
}
