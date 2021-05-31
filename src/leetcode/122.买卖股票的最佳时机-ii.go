/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit(prices []int) int {
	n := len(prices)
	var price int
	profit := 0
	for i := 1; i < n; i++ {
		price = prices[i] - prices[i-1]
		if price > 0 {
			profit += price
		}
	}
	return profit
}
// @lc code=end

