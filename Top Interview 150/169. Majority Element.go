package main

func majorityElement(nums []int) int {

	majority := 0
	res := 0

	for _, num := range nums {

		if majority == 0 {
			res = num
		}

		if num == res {
			majority++
		} else {
			majority--
		}

	}

	return res
}
