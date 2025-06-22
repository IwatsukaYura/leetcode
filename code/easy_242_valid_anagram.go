package code

/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mapS := make(map[rune]int)

	for _, valS := range s {
		mapS[valS]++
	}

	countLetters := len(mapS)

	for _, valT := range t {
		if _, ok := mapS[valT]; !ok {
			return false
		}
		mapS[valT]--

		if mapS[valT] < 0 {
			return false
		}

		if mapS[valT] == 0 {
			countLetters--
		}
	}
	return countLetters == 0

}

// @lc code=end
