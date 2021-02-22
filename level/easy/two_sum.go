package easy

func twoSum(nums []int, target int) []int {
	count = 0
	b := make([]int, 2)
	length := len(nums)
	for count < length {
		findNum := target - nums[count]
		i, ok := Find(nums, findNum)
		if ok != false {
			b[0] = count
			b[1] = i
			return b
		}
		count++
	}
	return b
}

func Find(a []int, findNum int) (int, bool) {
	for i := count + 1; i < len(a); i++ {
		if findNum == a[i] {
			return i, true
		}
	}
	return 0, false
}

func twoSum2(nums []int, target int) []int {
	m := map[int]int{}
	for i, n := range nums {
		if j, ok := m[target-n]; ok {
			return []int{j, i}
		}
		m[n] = i
	}
	return nil
}
