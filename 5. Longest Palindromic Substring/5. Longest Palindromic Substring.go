package main

import "strings"

func preprocess(s string) string {
	n := len(s)
	var newS strings.Builder
	newS.Grow(2*n + 1)
	newS.WriteByte('!')
	for i := 0; i < n; i++ {
		newS.WriteByte(s[i])
		newS.WriteByte('!')
	}
	return newS.String()
}

func longestPalindrome(s string) string {
	newS := preprocess(s)
	n := len(newS)
	m := make([]int, n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			k := i - (j - i)
			if k < 0 || newS[k] != newS[j] {
				break
			}
			m[i]++
		}
	}
	Max := 0
	index := 0
	for i, v := range m {
		if v > Max {
			Max = v
			index = i
		}
	}
	var res strings.Builder
	res.Grow(Max)
	for _, c := range newS[index-Max : index+Max+1] {
		if c == '!' {
			continue
		}
		res.WriteRune(c)
	}

	return res.String()

	//n := len(s)
	//P := make([]int, n)
	//C := 0
	//R := 0
	//for i := 0; i < n; i++ {
	//	mirror := 2*C - i
	//
	//	if i < R {
	//		if R-i < P[mirror] {
	//			P[i] = R - i
	//		} else {
	//			P[i] = P[mirror]
	//		}
	//	}
	//
	//	for i+P[i]+1 < n && i-P[i]-1 >= 0 && s[i+P[i]+1] == s[i-P[i]-1] {
	//		P[i] += 1
	//	}
	//
	//	if i+P[i] > R {
	//		C = i
	//		R = i + P[i]
	//	}
	//}

}
