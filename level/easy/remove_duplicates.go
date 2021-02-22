package easy

var count int

func removeDuplicates(nums []int) int {
	count = 0
	length := len(nums)
	for count < length {
		removeDuplicate(&nums)
		count++
	}
	return len(nums)
}

func removeDuplicate(a *[]int) {
	for i := count + 1; i < len(*a); i++ {
		if (*a)[count] == (*a)[i] {
			*a = append((*a)[:i], (*a)[i+1:]...)
			i = count
		}
	}
}
