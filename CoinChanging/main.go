package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Solution([]int{1}, 5))
	fmt.Println(Solution([]int{1, 3}, 5))
	fmt.Println(Solution([]int{1, 3, 4}, 5))
	fmt.Println(Solution([]int{1, 3, 4, 5}, 5))
	fmt.Println(Solution([]int{1}, 6))
	fmt.Println(Solution([]int{1, 3}, 6))
	fmt.Println(Solution([]int{1, 3, 4}, 6))
	fmt.Println(Solution([]int{1, 3, 4, 5}, 6))
	fmt.Println(Solution([]int{1, 3, 4, 5, 6}, 6))
}

func Solution(C []int, k int) int {
	n := len(C) + 1
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}

	for i := 0; i < n; i++ {
		dp[i][0] = 0
	}

	for i := 1; i <= k; i++ {
		dp[0][i] = math.MaxInt64
	}

	for i := 1; i < n; i++ {
		for j := 1; j <= k; j++ {
			Cmax := C[0]
			for idx := 0; idx < i; idx++ {
				if C[idx] > j {
					break
				}
				Cmax = C[idx]
			}

			dpij1 := dp[i][j-Cmax] + 1
			dpij2 := dp[i-1][j]
			if dpij1 <= dpij2 {
				dp[i][j] = dpij1
			} else {
				dp[i][j] = dpij2
			}
		}
	}

	//printDp(dp)

	return dp[n-1][k]
}

func printDp(dp [][]int) {
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[0]); j++ {
			fmt.Printf("%d ", dp[i][j])
		}
		fmt.Printf("\n")
	}
}
