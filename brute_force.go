package main

// countSubsetSumsLessThanOrEqual counts the subset sums that are <= k
func BruteForceKnapsack(weights []int, capacity int) int {
	n := len(weights)
	totalSubsets := 1 << n // 2^n subsets
	counter := 0

	// Iterate through all subsets
	for i := 0; i < totalSubsets; i++ {
		sum := 0

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
