package main

import "fmt"

func main() {
	testcases, err := readTestData("test_data.txt")
	if err == nil {
		for _, problem := range testcases {
			fmt.Println(NaiveDPKnapsack(problem.Weights, problem.Capacity))
			fmt.Println(BruteForceKnapsack(problem.Weights, problem.Capacity))
		}
	}

	// Calculate the result
	//fmt.Println(NaiveDPKnapsack(weight, W))
	//fmt.Println(bruteForceKnapsack(weight, W))
}
