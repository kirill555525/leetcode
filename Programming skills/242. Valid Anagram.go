package main

func isAnagram(s string, t string) bool {
	m := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	for j := 0; j < len(t); j++ {
		m[t[j]]--
		if m[t[j]] == 0 {
			delete(m, t[j])
		}
	}

	return len(m) == 0
}
