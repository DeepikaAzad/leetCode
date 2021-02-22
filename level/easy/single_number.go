package easy

func singleNumber(nums []int) int {
	m := map[int]int{}
	for _, n := range nums {
		if _, ok := m[n]; ok {
			return
		}
		m[n] = 1
	}
	return false
}
