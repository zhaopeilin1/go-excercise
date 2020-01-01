package leetcode88

func merge(nums1 []int, m int, nums2 []int, n int) {
	//给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。
	//初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
	//你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
	num1Index := (m - 1)
	num2Index := (n - 1)
	for i := n + m - 1; i >= 0; i-- {
		if num1Index >= 0 {
			if num2Index >= 0 {
				if nums1[num1Index] > nums2[num2Index] {
					nums1[i] = nums1[num1Index]
					num1Index--
				} else {
					nums1[i] = nums2[num2Index]
					num2Index--
				}
			}
		} else {
			if num2Index >= 0 {
				nums1[i] = nums2[num2Index]
				num2Index--
			}
		}
	}
}
