/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	n := len(prices)
	min, max := math.MaxInt32, 0
	for i := 0; i < n; i++ {
		if prices[i] < min {
			min = prices[i]
		}
		price := prices[i]-min
		if price > max {
			max = price
		}
	}
	return max
}

// @lc code=end

