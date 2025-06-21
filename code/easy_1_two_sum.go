package code
/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start

func twoSum(nums []int, target int) []int {
	resultMap := map[int]int{}
	for index, value := range nums {
		//value + want = target
		want := target - value
		if _, ok := resultMap[want]; ok != false {
			return []int{resultMap[want], index}
		}
		resultMap[value] = index
	}
	
	return nil
}
// @lc code=end

