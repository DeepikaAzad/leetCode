package easy

func containDuplicate(nums []int) bool {
	m := map[int]int{}
	for _, n := range nums {
		if _, ok := m[n]; ok {
			return ok
		}
		m[n] = 1
	}
	return false
}
