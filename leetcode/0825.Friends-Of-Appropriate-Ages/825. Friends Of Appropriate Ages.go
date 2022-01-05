package leetcocde

import "sort"

// 解法一
func numFriendRequests(ages []int) int {
	return 0
}

// 解法二 双指针 + 排序 O(n logn)
func numFriendRequests1(ages []int) int {
	sort.Ints(ages)
	left, right, res := 0, 0, 0
	for _, age := range ages {
		if age < 15 {
			continue
		}
		for ages[left]*2 <= age+14 {
			left++
		}
		for right+1 < len(ages) && ages[right+1] <= age {
			right++
		}
		res += right - left
	}
	return res
}

// 解法三 暴力解法 O(n^2)
func numFriendRequests2(ages []int) int {
	res, count := 0, [125]int{}
	for _, x := range ages {
		count[x]++
	}
	for i := 1; i <= 120; i++ {
		for j := 1; j <= 120; j++ {
			if j > i {
				continue
			}
			if (j-7)*2 <= i {
				continue
			}
			if j > 100 && i < 100 {
				continue
			}
			if i != j {
				res += count[i] * count[j]
			} else {
				res += count[i] * (count[j] - 1)
			}
		}
	}
	return res
}
