package main

import "fmt"

func climbStairs(n int) int {
	var ans = []int{1, 1}
	for i := 2; i <= n; i++ {
		ans = append(ans, ans[i-2]+ans[i-1])
	}
	return ans[n]
}

func main() {
	/*
		n: stairs total, can climb 1 or 2 steps for each time
		total possible climb ways match fibo array
	*/
	var n int = 10
	fmt.Println(climbStairs(n))
}
