package leet_0080

func RemoveDuplicates(nums []int) int {
	length := len(nums)
	if length < 3 {
		return length
	}
	end := 0
	cur := 1
	count := 1
	for cur < length {
		if nums[end] == nums[cur] {
			if count < 2 {
				end++
				count++
				nums[end] = nums[cur]
				cur++
				continue
			} else {
				for cur < length && nums[end] == nums[cur] {
					cur++
				}
			}
		}
		if cur < length {
			end++
			count = 1
			nums[end] = nums[cur]
		}
		cur++
	}
	return end+1
}
