package main

import "strconv"

func reverse(x int) int {
	flag := true
	if x < 0 {
		flag = false
		x = -x
	}
	s := make([]byte, 0, 10)
	for i := 0; x != 0; i++ {
		s = append(s, byte(48+x%10))
		x /= 10
	}
	str := string(s)
	if len(str) == 10 && ((str > "2147483647" && flag) || (str > "2147483648" && !flag)) {
		return 0
	}

	res, _ := strconv.Atoi(str)
	if !flag {
		res = -res
	}
	return res

}
