package main

import (
	"strings"
)

func mergeAlternately(word1 string, word2 string) string {
	var s strings.Builder

	for i := 0; i < len(word1) || i < len(word2); i++ {
		if i < len(word1) {
			s.WriteByte(word1[i])
		}

		if i < len(word2) {
			s.WriteByte(word2[i])
		}
	}

	return s.String()

}
