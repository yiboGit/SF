package main

func main() {

}

// O(log n)可以想到是 二分法 这里的二分法 有点特殊
// 因为旋转的存在，只能保证二分之后的一边是有序的,那我们只查看有序的一边
func search(nums []int, target int) int {

	start, end, mid := 0, len(nums)-1, 0

	for start <= end {
		mid = (start + end) / 2
		// 找到
		if nums[mid] == target {
			return mid
		}
		//找到二分之后有顺序的一边
		// 看右边是否满足有序条件
		if nums[mid] < nums[end] {
			//满足条件，右边有序
			// 查看target 是否在此有序区间内
			if target <= nums[end] && target > nums[mid] {
				// 在 在有顺序一边
				start = mid + 1
			} else {
				// 不在；在无顺序一边
				end = mid - 1
			}
		} else {
			//不满足条件，说明左边有序
			// 查看target 是否在此有序区间内
			if target >= nums[start] && target < nums[mid] {
				// 在；end = 中位-1  在有顺序一边 ，缩小二分范围
				end = mid - 1
			} else {
				// 不在； start = 中位+1 在无顺序一边
				start = mid + 1
			}
		}

	}
	return -1
}
