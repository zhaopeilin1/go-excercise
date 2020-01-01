package leetcode88

func merge(nums1 []int, m int, nums2 []int, n int) {
	//给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。
	//初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
	//你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
	//nums1 0 从nums1和nums2中的 0取。
	num2Index = 0
	num1Index = 0
	//存放在num2中的num1的元素index
	num1Index2 = -1

	for i := 0; i < n+m; i++ {
		if num1Index2 < 0 {
			//比较num1,num2index
			if nums2[num2Index] < nums1[i] {
				nums1[i], nums2[num2Index] = nums2[num2Index], nums1[i]
				num1Index2++
				num2Index++
			} else {
			}
		} else {
			//比较index2和 num2index
			if nums2[num2Index] < nums2[num1Index2] {
				nums1[i], nums2[num2Index] = nums2[num2Index], nums1[i]
				num1Index2++
				num2Index++
			} else {

			}
		}
	}
}
