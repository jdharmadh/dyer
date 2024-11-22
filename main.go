package main

import (
	"fmt"
)

func main() {
	generateTestCases("test_data.txt")
	testcases, err := readTestData("test_data.txt")
	if err == nil {
		for _, problem := range testcases {
			if NaiveDPKnapsack(problem.Weights, problem.Capacity) != problem.Answer {
				fmt.Println(problem)
			}
			//fmt.Println(BruteForceKnapsack(problem.Weights, problem.Capacity))
		}
	}

	// Calculate the result
	//fmt.Println(NaiveDPKnapsack(weight, W))
	//fmt.Println(bruteForceKnapsack(weight, W))
}
