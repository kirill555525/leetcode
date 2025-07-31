package main

func subarrayBitwiseORs(arr []int) int {
	m := make(map[int]struct{})
	cur := map[int]struct{} {
		0: struct{}{},
	}
	for _, num := range arr {
		temp := map[int]struct{}{}
		for key := range cur {
			temp[num | key] = struct{}{}
			temp[num] = struct{}{}
		}
		cur = temp
		for key := range temp {
			m[key] = struct{}{}
		}
	}

	return len(m)

}