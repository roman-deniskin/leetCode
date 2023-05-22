package climbing_stairs

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	firstNum := 1
	secondNum := 1
	for i := 1; i <= n; i++ {
		firstNum, secondNum = secondNum, firstNum+secondNum
	}
	return firstNum
}
