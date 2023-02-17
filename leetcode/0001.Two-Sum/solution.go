package leetcode

func twoSumSolution(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for k, v := range nums {
		if n, ok := m[target-v]; ok {
			return []int{n, k}
		}
		m[v] = k
	}

	return nil
}
