package main

func findWordsContaining(words []string, x byte) []int {
	var res []int
	for i, s := range words {
		for j := range s {
			if s[j] == x {
				res = append(res, i)
				break
			}
		}
	}
	return res
}
