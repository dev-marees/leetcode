package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	tokens := []string{"2", "1", "+", "3", "*"}
	tokens2 := []string{"4", "13", "5", "/", "+"}
	println("evalRPN ::", evalRPN(tokens))
	println("evalRPN ::", evalRPN(tokens2))
}

func evalRPN(tokens []string) int {
	stackArray := []int{}
	tempArray := []int{}
	for i := 0; i < len(tokens); i++ {
		intVal, err := strconv.Atoi(tokens[i])
		fmt.Println("intVal::", intVal)
		if err != nil && strings.Contains(err.Error(), "invalid syntax") {
			switch tokens[i] {
			case "+":
			case "-":
			case "/":
			case "*":
			}
		} else {
			stackArray = append(stackArray, intVal)
		}
	}
	return tempArray[0]
}
