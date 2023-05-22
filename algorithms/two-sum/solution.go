package two_sum

func twoSum(nums []int, target int) []int {
	var res []int
	for index1, num1 := range nums {
		for index2, num2 := range nums {
			if index1 == index2 {
				continue
			}
			if num1+num2 == target {
				return []int{index1, index2}
			}
		}
	}
	return res
}
