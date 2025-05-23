package main

func lengthOfLongestSubstring(s string) int {

	m := make(map[byte]int)
	Max := 0
	left := 0

	for right := 0; right < len(s); right++ {
		c := s[right]
		if v, ok := m[c]; ok && v >= left {
			left = v + 1
		}
		m[c] = right

		if Max < right-left+1 {
			Max = right - left + 1
		}

	}

	return Max

}
