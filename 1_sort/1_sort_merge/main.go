package main

func main() {

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	oneIndex := m - 1 //nums1
	twoIndex := n - 1 //nums2
	inputIndex := m + n - 1

	for twoIndex >= 0 {
		if oneIndex < 0 || nums1[oneIndex] < nums2[twoIndex] {
			nums1[inputIndex] = nums2[twoIndex]
			twoIndex--
		} else {
			nums1[inputIndex] = nums1[oneIndex]
			oneIndex--

		}
		inputIndex--
	}
}
