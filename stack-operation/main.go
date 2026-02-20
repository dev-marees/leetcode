package main

import "fmt"

func main() {
	nums := []int{1, 3}
	nums1 := []int{1, 2, 3}
	fmt.Println("build Array ::", buildArray(nums, 3))
	fmt.Println("build Array ::", buildArray(nums1, 3))
}

func buildArray(target []int, n int) []string {
	result := []string{}
	j := 0

	for x := 1; x <= n && j < len(target); x++ {
		result = append(result, "Push")
		if x == target[j] {
			j++
		} else {
			result = append(result, "Pop")
		}
	}

	return result
}
