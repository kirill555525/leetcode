package main

import (
	"fmt"
)

func lengthAfterTransformations(s string, t int) int {
	arr := [26]int{}
	for i := 0; i < len(s); i++ {
		arr[s[i]-'a']++
	}

	for i := 0; i < t; i++ {
		z := arr[25]
		for j := 25; j > 0; j-- {
			arr[j] = arr[j-1]

		}
		arr[0] = z
		arr[1] = (arr[1] + z) % (1_000_000_000 + 7)

	}

	total := 0

	for _, v := range arr {
		total = (total + v) % (1_000_000_000 + 7)
	}

	return total

}

func main() {
	s := "jqktcurgdvlibczdsvnsg"
	t := 7517
	fmt.Println(lengthAfterTransformations(s, t))

}
