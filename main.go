package main

import "fmt"

func main() {
	// Weights and knapsack capacity
	weight := []int{10, 20, 30, 40, 82}
	W := 120
	n := len(weight)

	// Initialize DP table with -1 (indicating uncomputed values)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, W+2)
		for j := 0; j <= W+1; j++ {
			dp[i][j] = -1
		}
	}

	// Base case: 1 way to achieve 0 capacity with 0 items
	dp[0][0] = 1

	// Calculate the result
	result := ks(n, W, dp, weight)
	fmt.Println(result)
}
