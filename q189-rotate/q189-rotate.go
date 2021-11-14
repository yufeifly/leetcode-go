package q189_rotate

func rotate(nums []int, k int) {
	k = k % len(nums)

	length := len(nums)

	nums = append(nums[length-k:], nums[:length-k]...)
}
