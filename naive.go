package main

// Function to calculate the number of ways to fill the knapsack
func ks(j uint, w uint, dp [][]uint, weight []uint, visited [][]uint) uint {
	// Base case: if already computed, return the result
	if visited[j][w] == 1 {
		return uint(dp[j][w])
	}

	visited[j][w] = 1

	// Base case: if no items left (j == 0), return 1
	if j == 0 {
		dp[j][w] = 1
		return dp[j][w]
	}

	// If the weight of the current item is less than or equal to the capacity,
	// either include it or exclude it.
	if weight[j-1] <= w {
		dp[j][w] = ks(j-1, w, dp, weight, visited) + ks(j-1, w-weight[j-1], dp, weight, visited)
	} else {
		dp[j][w] = ks(j-1, w, dp, weight, visited)
	}

	return dp[j][w]
}

func NaiveDPKnapsack(weights []uint, capacity uint) uint {
	n := uint(len(weights))
	// Initialize DP table with -1 (indicating uncomputed values)
	dp := make([][]uint, n+1)
	visited := make([][]uint, n+1)
	for i := 0; i <= int(n); i++ {
		dp[i] = make([]uint, capacity+2)
		visited[i] = make([]uint, capacity+2)
		for j := 0; j <= int(capacity)+1; j++ {
			dp[i][j] = 0
			visited[i][j] = 0
		}
	}

	// Base case: 1 way to achieve 0 capacity with 0 items
	dp[0][0] = 1
	return ks(n, capacity, dp, weights, visited)
}

func SpaceEfficientKnapsack(weights []uint, capacity uint) uint {
	n := len(weights)
	if n == 0 {
		return 0
	}

	// We only need two rows: previous and current
	dp := make([][]uint, 2)
	dp[0] = make([]uint, capacity+1)
	dp[1] = make([]uint, capacity+1)

	// Base case: one way to fill a knapsack of any capacity with no items
	for w := uint(0); w <= capacity; w++ {
		dp[0][w] = 1
	}

	// Fill the dp table bottom up, considering each item
	for i := 1; i <= n; i++ {
		curr := i & 1
		prev := (i - 1) & 1

		for w := uint(0); w <= capacity; w++ {
			// Start with ways without current item
			dp[curr][w] = dp[prev][w]

			// If we can include the current item, add those possibilities
			if w >= weights[i-1] {
				dp[curr][w] += dp[prev][w-weights[i-1]]
			}
		}
	}

	// Return the final result
	return dp[n&1][capacity]
}
