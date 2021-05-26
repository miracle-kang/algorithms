/*
 * @lc app=leetcode.cn id=690 lang=golang
 *
 * [690] 员工的重要性
 */

// @lc code=start
/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

func getImportance(employees []*Employee, id int) int {
	var employee *Employee = nil
	for _, e := range employees {
		if e.Id == id {
			employee = e
			break
		}
	}
	if employee == nil {
		return 0
	}
	importance := employee.Importance
	for _, subId := range employee.Subordinates {
		importance += getImportance(employees, subId)
	}
	return importance
}

// @lc code=end

