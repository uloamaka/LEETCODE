func longestOnes(nums []int, k int) int {
    var maxOnes, l, zeroCnt int
    
    for r := 0; r < len(nums); r++ {
        if nums[r] == 0 {
            zeroCnt++
        }

        for zeroCnt > k {
            if nums[l] == 0 {
                zeroCnt--
            }
            l++
        }
        
        winSize := r-l+1

        if winSize > maxOnes {
            maxOnes = winSize
        }
    } 
    return maxOnes 
}