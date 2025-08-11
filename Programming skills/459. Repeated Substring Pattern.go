package main

import "strings"

func repeatedSubstringPattern(s string) bool {
	str := s + s
	return strings.Contains(str[1:len(str)-1], s)
}
