package code
/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start

func twoSum(nums []int, target int) []int {
	result := []int{}
    for i:=0; i<len(nums); i++{
		for j:=i+1;j<len(nums); j++{
			addNum := nums[i] + nums[j]
			if (addNum == target){
				result = []int{i,j}
			}
		}
	}
	return result
}
// @lc code=end

