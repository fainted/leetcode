package main

import "log"

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	count, odd := (len(nums1)+len(nums2))>>1, (len(nums1)+len(nums2))&1
	if odd == 1 || count < len(nums1)+len(nums2) {
		// case odd: median = sorted(nums1, nums2)[ceil(count)];
		// case even: median = avg(sorted(nums1, nums2)[count, count+1])
		count++
	}

	var val1, val2 int
	for i1, i2, c := 0, 0, 0; c < count; c++ {
		if i1 < len(nums1) && i2 < len(nums2) {
			if nums1[i1] <= nums2[i2] {
				val1, val2 = val2, nums1[i1]
				i1++
			} else {
				val1, val2 = val2, nums2[i2]
				i2++
			}
		} else if i1 < len(nums1) {
			val1, val2 = val2, nums1[i1]
			i1++
		} else {
			val1, val2 = val2, nums2[i2]
			i2++
		}
	}

	if odd == 1 {
		return float64(val2)
	}
	return float64(val1+val2) / 2
}

func _test() {
	log.Print(findMedianSortedArrays([]int{}, []int{}))                  // 0
	log.Print(findMedianSortedArrays([]int{}, []int{1}))                 // 1
	log.Print(findMedianSortedArrays([]int{1, 3, 5}, []int{2, 4, 6}))    // 3.5
	log.Print(findMedianSortedArrays([]int{1, 3, 5, 7}, []int{2, 4, 6})) // 4
	log.Print(findMedianSortedArrays([]int{1, 3, 5}, []int{2, 4, 6, 8})) // 4
}
