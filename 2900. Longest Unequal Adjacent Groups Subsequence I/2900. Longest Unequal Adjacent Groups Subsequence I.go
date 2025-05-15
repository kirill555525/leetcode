package main

func getLongestSubsequence(words []string, groups []int) []string {
	output := make([]string, 0, len(words))
	n := len(groups)
	output = append(output, words[0])
	for i := 1; i < n; i++ {
		if groups[i] != groups[i-1] {
			output = append(output, words[i])
		}
	}
	return output
}
