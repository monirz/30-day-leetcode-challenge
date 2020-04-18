package main

func minPathSum(grid [][]int) int {

	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	dp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[0]))
	}

	m := len(grid)
	n := len(grid[0])

	dp[0][0] = grid[0][0]

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 && j > 0 {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j] //update the minimum
			}
			if i == 0 && j != 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j] //move to right
			}
			if j == 0 && i != 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j] // move down
			}
		}
	}

	return dp[m-1][n-1]

}
