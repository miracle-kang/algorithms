/*
 * @lc app=leetcode.cn id=1744 lang=golang
 *
 * [1744] 你能在你最喜欢的那天吃到你最喜欢的糖果吗？
 */

// @lc code=start
func canEat(candiesCount []int, queries [][]int) []bool {
	ans := make([]bool, len(queries))
	cnts := make([]int, len(candiesCount)+1)
	for i, cnt := range candiesCount {
		cnts[i+1] = cnts[i] + cnt
	}

	for i, query := range queries {
		day := query[1] + 1
		maxEat := day * query[2]
		ans[i] = cnts[query[0]+1] >= day && maxEat > cnts[query[0]]
	}
	return ans
}

// @lc code=end

