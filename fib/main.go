package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(coinChange([]int{1, 2, 5, 4}, 29))
}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 0
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func coinChange(coins []int, amount int) int {
	final := math.MaxInt64
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	for _, coin := range coins {
		subRes := coinChange(coins, amount-coin)
		if subRes == -1 {
			continue
		}
		final = int(math.Min(float64(final), float64(subRes+1)))
	}
	return final
}
