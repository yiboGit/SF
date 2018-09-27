package main

func main() {

}

// 使用二分法解决问题
func searchRange(nums []int, target int) []int {

	head, mid, end := 0, 0, len(nums)-1

	for head <= end {
		mid = (head + end) / 2
		if nums[mid] > target {
			//关注点 1 更新end,head索引值
			end = mid - 1
		} else if nums[mid] < target {
			head = mid + 1
		} else {
			// 关注点2 确定区间之后缩小范围
			if target > nums[head] {
				head++
			}
			if target < nums[end] {
				end--
			}
			if target == nums[head] && target == nums[end] {
				return []int{head, end}
			}
		}
	}
	return []int{-1, -1}
}
