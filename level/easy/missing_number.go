func missingNumber(nums []int) int {
    sort.Ints(nums)
    num, i := 0,0
    for i=0; i < len(nums); i++ {
        if i != nums[i] {
            num = i
            return num
        }
    }
    
    num = i
    return num
}