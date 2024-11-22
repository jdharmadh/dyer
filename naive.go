package main

// Function to calculate the number of ways to fill the knapsack
func ks(j, w int, dp [][]int, weight []int) int {
	// Base case: if already computed, return the result
	if dp[j][w] != -1 {
		return dp[j][w]
	}

	// Base case: if no items left (j == 0), return 1
	if j == 0 {
		dp[j][w] = 1
		return dp[j][w]
	}

	// If the weight of the current item is less than or equal to the capacity,
	// either include it or exclude it.
	if weight[j-1] <= w {
		dp[j][w] = ks(j-1, w, dp, weight) + ks(j-1, w-weight[j-1], dp, weight)
	} else {
		dp[j][w] = ks(j-1, w, dp, weight)
	}

	return dp[j][w]
}

func NaiveDPKnapsack(weights []int, capacity int) int {
	n := len(weights)

	// Initialize DP table with -1 (indicating uncomputed values)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, capacity+2)
		for j := 0; j <= capacity+1; j++ {
			dp[i][j] = -1
		}
	}

	// Base case: 1 way to achieve 0 capacity with 0 items
	dp[0][0] = 1
	return ks(n, capacity, dp, weights)
}
