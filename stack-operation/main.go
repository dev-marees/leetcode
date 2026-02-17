package main

import "fmt"

type BuildArray interface {
	push()
	pop()
	peak() int
}

type Stack struct {
	input  []int
	target int
}

func main() {
	nums := []int{1, 3}
	nums1 := []int{1, 2, 3}
	fmt.Println("build Array ::", buildArray(nums, 3))
	fmt.Println("build Array ::", buildArray(nums1, 3))
}

func (s Stack) push() {
	for i := 1; i <= s.target; i++ {
		if len(s.input) > 2 {
			//Find peak
			peakElement := s.peak()
			if s.target % peakElement == 0 {
				
			}
		}
		s.input = append(s.input, i)
	}
}

func (s Stack) peak() int {
	return s.input[len(s.input)-1]
}

func buildArray(nums []int, target int) []string {
	result := []string{}
	intArray := []int{}

	for i := 1; i <= target; i++ {
		intArray = append(intArray, 1)
	}

	return result
}
