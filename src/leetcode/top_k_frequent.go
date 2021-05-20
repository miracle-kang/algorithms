package leetcode

import "sort"

// 692. 前K个高频单词
// 给一非空的单词列表，返回前 k 个出现次数最多的单词。
// 返回的答案应该按单词出现频率由高到低排序。如果不同的单词有相同出现频率，按字母顺序排序。
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
