package exam

import (
	"math"
)

func Ft_min_window(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	charCount := make(map[rune]int)
	for _, char := range t {
		charCount[char]++
	}

	left, right := 0, 0
	required := len(charCount)
	formed := 0
	windowCounts := make(map[rune]int)

	minLen := math.MaxInt32
	minLeft, minRight := 0, 0

	for right < len(s) {
		char := rune(s[right])
		windowCounts[char]++

		if count, exists := charCount[char]; exists && windowCounts[char] == count {
			formed++
		}

		for left <= right && formed == required {
			char = rune(s[left])

			if right-left+1 < minLen {
				minLen = right - left + 1
				minLeft = left
				minRight = right
			}

			windowCounts[char]--
			if count, exists := charCount[char]; exists && windowCounts[char] < count {
				formed--
			}
			left++
		}
		right++
	}

	if minLen == math.MaxInt32 {
		return ""
	}
	return s[minLeft : minRight+1]
}
