package hard

import "neekity.com/leetcode/src/common"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var res float64
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		return float64(findKthNumber(nums1, nums2, totalLength/2+1))
	} else {
		res = float64(findKthNumber(nums1, nums2, totalLength/2+1)+findKthNumber(nums1, nums2, totalLength/2)) / 2
	}
	return res
}

func findKthNumber(nums1 []int, nums2 []int, kth int) int {
	lenShort, lenLong := len(nums1), len(nums2)
	if lenShort > lenLong {
		return findKthNumber(nums2, nums1, kth)
	}
	if lenShort == 0 {
		return nums2[kth-1]
	}

	if kth <= lenShort {
		return findMidNumber(nums1, nums2, 0, 0, kth, kth)
	}

	if kth > lenLong {
		if nums2[kth-lenShort-1] >= nums1[lenShort-1] {
			return nums2[kth-lenShort-1]
		}

		if nums1[kth-lenLong-1] >= nums2[lenLong-1] {
			return nums1[kth-lenLong-1]
		}
		return findMidNumber(nums1, nums2, kth-lenLong, kth-lenShort, lenShort, lenLong)
	}

	if nums1[lenShort-1] <= nums2[kth-lenShort-1] {
		return nums2[kth-lenShort-1]
	}

	return findMidNumber(nums1, nums2, 0, kth-lenShort, lenShort, kth)
}

func findMidNumber(nums1 []int, nums2 []int, l1 int, l2 int, r1 int, r2 int) int {
	mid1, mid2, offset := 0, 0, 0

	for {
		if (r1-l1)%2 == 0 {
			mid1, mid2 = (r1+l1)/2-1, (r2+l2)/2-1
		} else {
			mid1, mid2 = (r1+l1)/2, (r2+l2)/2
		}
		if l1 >= r1-1 {
			break
		}
		offset = 1 - (r1-l1)%2

		if nums1[mid1] < nums2[mid2] {
			l1 = mid1 + offset
			r2 = mid2 + 1
		} else if nums1[mid1] > nums2[mid2] {
			r1 = mid1 + 1
			l2 = mid2 + offset
		} else {
			return nums1[mid1]
		}
	}

	return common.Min(nums1[l1], nums2[l2])
}
