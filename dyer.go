package main

import (
	"math"
	"math/rand"
)

func SampleWeightFromTable(table [][]uint, weights []uint) uint {
	n := uint(len(weights))
	var T uint
	T = 0
	j := n
	k := n * n
	for j > 0 {
		probability := float64(table[j-1][k-weights[j-1]]) / float64(table[j][k])
		r := rand.Float64()
		if r < probability {
			T += weights[j-1]
			k -= weights[j-1]
		}
		j--
	}
	return T
}

func ApproximateSolution(weights []uint, capacity uint, err float64, conf float64) uint {
	n := uint(len(weights))
	N := int(4 * float64(n) * math.Log(2/conf) / (err * err))
	round_factor := float64(n*n) / float64(capacity)
	new_weights := make([]uint, n)
	for i := 0; i < len(weights); i++ {
		new_weights[i] = uint(math.Floor(float64(weights[i]) * round_factor))
	}
	//print before weights and after weights
	table := KnapsackTable(new_weights, n*n)

	// take N samples and count the number of times its a valid solution
	// fmt.Println(N)
	count := 0
	for i := 0; i < N; i++ {
		T := SampleWeightFromTable(table, new_weights)
		// fmt.Println(T)
		if T <= capacity {
			count++
		}
	}

	//randomly sample from the table
	return uint(N)
}
