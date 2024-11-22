package main

import (
	"fmt"
	"math"
	"math/rand"
)

// countSubsetSumsLessThanOrEqual counts the subset sums that are <= k
func BruteForceKnapsack(weights []uint, capacity uint) uint {
	n := len(weights)
	totalSubsets := 1 << n // 2^n subsets
	var counter uint = 0

	// Iterate through all subsets
	for i := 0; i < totalSubsets; i++ {
		var sum uint = 0

		// Calculate the sum of the current subset
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				sum += weights[j]
			}
		}

		// Check if the sum is less than or equal to k
		if sum <= capacity {
			counter++
		}
	}

	return counter
}

func generateTestCases(filename string) {

	var testCases []testCase

	for i := 0; i < 1000; i++ {
		// Generate a random number of weights (length between 5 and 10)
		length := uint(rand.Intn(10) + 30) // Random between 5 and 62

		// Generate random weights (values between 1 and 100)
		weights := make([]uint, length)
		for j := range weights {
			weights[j] = uint(rand.Intn(1000) + 1)
		}

		// Determine the capacity
		maxWeight := uint(0)
		for _, w := range weights {
			if w > maxWeight {
				maxWeight = w
			}
		}
		scale := 1 + rand.Float64()*0.7
		capacity := uint(rand.Intn(int(math.Pow(float64(maxWeight), scale))-int(maxWeight)+1) + int(maxWeight))

		// Calculate the answer using BruteForceKnapsack
		answer := NaiveDPKnapsack(weights, capacity)

		// Add the test case
		testCases = append(testCases, testCase{
			Capacity: capacity,
			Weights:  weights,
			Answer:   answer,
		})
	}

	// Write the test cases to "test_data.txt"
	err := writeTestData(filename, testCases)
	if err != nil {
		fmt.Println("Error writing test data:", err)
	} else {
		fmt.Println("Test data written successfully to " + filename)
	}
}
