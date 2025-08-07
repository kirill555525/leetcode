package main

func findTheDifference(s string, t string) byte {
	m := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		m[t[i]]--
		if m[t[i]] == -1 {
			return t[i]
		}
	}

	return 0
}
