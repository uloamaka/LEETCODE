func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)

    closest := math.MaxInt
    ans := 0
    for i := 0; i < len(nums)-2; i++ {
        
        l, r := i+1, len(nums)-1
        for l<r {
            sum := nums[i]+nums[l]+nums[r]
            diff := target - sum
            
            diffAbs := abs(diff)
            if diffAbs == 0 {
                return sum
            }

            if diffAbs < closest {
                closest = diffAbs
                ans = sum
            }

            if diff > 0{
                l++
            } else {
                r--
            }
        }
    }
    return ans
}

func abs(num int) int {
    if num < 0 {
        return -num
    }
    return num
}