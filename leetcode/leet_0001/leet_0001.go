package leet_0001

func TwoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for i, num := range (nums) {
		if index, ok := indexMap[target-num]; ok {
			return []int{index, i}
		} else {
			indexMap[num] = i
		}
	}
	return nil
}
