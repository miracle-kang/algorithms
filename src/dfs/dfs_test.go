package dfs_test

import (
	"algorithms/src/dfs"
	"testing"
)

func TestCourseQuery(t *testing.T) {

	// 4
	// [[2,3],[2,1],[0,3],[0,1]]
	// [[0,1],[0,3],[2,3],[3,0],[2,0],[0,2]]
	numCourses := 4
	prerequisites := [][]int{{2, 3}, {2, 1}, {0, 3}, {0, 1}}
	queries := [][]int{{0, 1}, {0, 3}, {2, 3}, {3, 0}, {2, 0}, {0, 2}}

	res := dfs.CheckIfPrerequisite(numCourses, prerequisites, queries)
	t.Log(res)

	// 5
	// [[4,3],[4,1],[4,0],[3,2],[3,1],[3,0],[2,1],[2,0],[1,0]]
	// [[1,4],[4,2],[0,1],[4,0],[0,2],[1,3],[0,1]]
	numCourses = 5
	prerequisites = [][]int{{4, 3}, {4, 1}, {4, 0}, {3, 2}, {3, 1}, {3, 0}, {2, 1}, {2, 0}, {1, 0}}
	queries = [][]int{{1, 4}, {4, 2}, {0, 1}, {4, 0}, {0, 2}, {1, 3}, {0, 1}}

	res = dfs.CheckIfPrerequisite(numCourses, prerequisites, queries)
	t.Log(res)
}
