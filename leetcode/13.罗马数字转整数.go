/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	count := 0
	for i := 0; i <= len(s)-1; {
		switch string(s[i]) {
		case "I":
			if len(s)-1 > i && string(s[i+1]) == "V" {
				count += 4
				i += 2
				continue
			}
			if len(s)-1 > i && string(s[i+1]) == "X" {
				count += 9
				i += 2
				continue
			}
			fallthrough
		case "X":
			if len(s)-1 > i && string(s[i+1]) == "L" {
				count += 40
				i += 2
				continue
			}
			if len(s)-1 > i && string(s[i+1]) == "C" {
				count += 90
				i += 2
				continue
			}
			fallthrough
		case "C":
			if len(s)-1 > i && string(s[i+1]) == "D" {
				count += 400
				i += 2
				continue
			}
			if len(s)-1 > i && string(s[i+1]) == "M" {
				count += 900
				i += 2
				continue
			}
			fallthrough
		default:
			count += romanMap[string(s[i])]
			i++
		}

	}
	return count
}

// @lc code=end

