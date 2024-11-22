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
