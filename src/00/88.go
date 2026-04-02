package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	n--
	m--
	for i := n + m + 1; i >= 0; i-- {
		if n < 0 {
			break
		}

		if m < 0 || nums1[m] < nums2[n] {
			nums1[i] = nums2[n]
			n--
		} else {
			nums1[i] = nums1[m]
			m--
		}
	}
}
