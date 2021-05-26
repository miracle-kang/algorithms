/*
 * @lc app=leetcode.cn id=692 lang=golang
 *
 * [692] 前K个高频单词
 */

// @lc code=start
func topKFrequent(words []string, k int) []string {
	cntMap := make(map[string]int)
	for _, word := range words {
		cntMap[word]++
	}

	keys := make([]string, len(cntMap))
	for word := range cntMap {
		keys = append(keys, word)
	}
	sort.Slice(keys, func(i, j int) bool {
		return cntMap[keys[i]] > cntMap[keys[j]] || cntMap[keys[i]] == cntMap[keys[j]] && keys[i] < keys[j]
	})
	return keys[0:k]
}

// @lc code=end

