/*
 * @lc app=leetcode.cn id=516 lang=golang
 *
 * [516] 最长回文子序列
 */

// @lc code=start
func longestPalindromeSubseq(s string) int {
    n := len(s)
    dp := make([][]int, n)
    for i := range(dp) {
        dp[i] = make([]int, n)
    }
    for i := n-1; i >= 0; i-- {
        dp[i][i] = 1
        for j := i+1; j < n; j++ {
            if s[i] == s[j] {
                dp[i][j] = dp[i+1][j-1] + 2
            } else {
                dp[i][j] = max(dp[i][j-1], dp[i+1][j])
            }
        }
    }
    return dp[0][n-1]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
// @lc code=end

