func threeSum(nums []int) [][]int {
    sort.Ints(nums)

    ans := [][]int{}
    for i := 0; i<len(nums)-2; i++ {
        if i > 0 && nums[i-1] == nums[i] {
            continue
        }

        l, r := i+1, len(nums)-1
        for l < r {
            sum := nums[i]+nums[l]+nums[r]

            switch {
            case sum < 0:
                l++
            case sum > 0:
                r--
            default:
                ans = append(ans, []int{nums[i],nums[l],nums[r]})

                for l < r && nums[l]==nums[l+1] {
                    l++
                }

                for l < r && nums[r]==nums[r-1] {
                    r--
                }

                l++
                r--
            } 
        }
    }
    return ans
}
// [-4,-4,-1,-1,0,1,2]