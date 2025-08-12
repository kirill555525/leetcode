package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	m--
	n--
	for i := m + n + 1; i >= 0; i-- {
		switch {
		case m == -1:
			nums1[i] = nums2[n]
			n--
		case n == -1:
			nums1[i] = nums1[m]
			m--
		case nums1[m] >= nums2[n]:
			nums1[i] = nums1[m]
			m--
		case nums1[m] < nums2[n]:
			nums1[i] = nums2[n]
			n--
		}

	}

}
