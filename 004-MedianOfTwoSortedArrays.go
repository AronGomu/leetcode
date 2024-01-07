package main

// import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// fmt.Printf("nums1 : %v\nnums2 : %v\n", nums1, nums2)

	index1 := 0
	index2 := 0

	merged := []int{}

	for {

		// fmt.Printf("index: %v & %v\n", index1, index2)
		// fmt.Printf("%v >= %v\n", nums1[index1], nums2[index2])

		if index1 > len(nums1)-1 {
			merged = append(merged, nums2[index2])
			if index2 < len(nums2) {
				index2 += 1
			}
		} else if index2 > len(nums2)-1 {
			merged = append(merged, nums1[index1])
			if index1 < len(nums1) {
				index1 += 1
			}
		} else if nums1[index1] <= nums2[index2] {
			merged = append(merged, nums1[index1])
			if index1 < len(nums1) {
				index1 += 1
			}
		} else if nums1[index1] > nums2[index2] {
			merged = append(merged, nums2[index2])
			if index2 < len(nums2) {
				index2 += 1
			}
		}

		// fmt.Printf("merged: %v\n", merged)

		if index1 > len(nums1)-1 && index2 > len(nums2)-1 {
			break
		}

	}

	var median float64
	l := len(merged)
	if l == 0 {
		return 0
	} else if l%2 == 0 {
		median = float64(merged[l/2-1]+merged[l/2]) / 2
	} else {
		median = float64(merged[l/2])
	}

	// fmt.Printf("median : %v\n\n", median)
	return median
}
