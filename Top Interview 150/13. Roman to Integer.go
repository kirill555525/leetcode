package main

func romanToInt(s string) int {
	total := 0

	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := 0; i < len(s); i++ {

		if i != len(s)-1 && m[s[i+1]] > m[s[i]] {
			total -= m[s[i]]
		} else {
			total += m[s[i]]
		}

	}

	return total

}
