package dfs

import "strconv"

func CheckIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	mp := make(map[int][]int, numCourses)
	for _, pre := range prerequisites {
		mp[pre[1]] = append(mp[pre[1]], pre[0])
	}

	cache := make(map[string]bool)
	res := make([]bool, len(queries))
	for i, query := range queries {
		res[i] = dfs(mp, cache, query[0], query[1])
	}
	return res
}

func dfs(mp map[int][]int, cache map[string]bool, a, b int) bool {
	if mp[b] == nil {
		return false
	}

	key := strconv.Itoa(a) + "-" + strconv.Itoa(b)
	if v, ok := cache[key]; ok {
		return v
	}

	for _, v := range mp[b] {
		k := strconv.Itoa(a) + "-" + strconv.Itoa(v)
		if val, ok := cache[k]; ok {
			if val {
				return true
			}
			continue
		}

		val := v == a || dfs(mp, cache, a, v)
		cache[strconv.Itoa(a)+"-"+strconv.Itoa(v)] = val
		if val {
			return true
		}
	}
	cache[key] = false
	return false
}
