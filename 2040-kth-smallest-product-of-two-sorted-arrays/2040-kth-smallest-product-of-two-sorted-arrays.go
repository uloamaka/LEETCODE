func kthSmallestProduct(nums1 []int, nums2 []int, k int64) int64 {
     // Step 1: Find search bounds
    n1 := len(nums1)
    n2 := len(nums2)
    minProd := min(nums1[0]*nums2[0], nums1[0]*nums2[n2-1], 
                  nums1[n1-1]*nums2[0], nums1[n1-1]*nums2[n2-1])
    maxProd := max(nums1[0]*nums2[0], nums1[0]*nums2[n2-1], 
                  nums1[n1-1]*nums2[0], nums1[n1-1]*nums2[n2-1])
    
    left := minProd
    right := maxProd
    
    // Step 2: Binary search on answer
    for left < right{
        mid := left + (right - left) / 2
        count := 0
        
        // Step 3: Count products <= mid
        for i := 0; i < len(nums1); i++ {
            if nums1[i] == 0 {
                if mid >= 0 {
                    count += len(nums2)
                }
            } else if nums1[i] > 0 {
                target := int(math.Floor(float64(mid) / float64(nums1[i])))
                count += binarySearchLE(nums2, target)
            } else {
                target := int(math.Ceil(float64(mid) / float64(nums1[i])))
                count += binarySearchGE(nums2, target)
            }
        }
        // Step 4: Update search bounds
        if count >= int(k) {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return int64(left)
}
   
func binarySearchLE(arr[]int, target int) int {
    left := 0
    right := len(arr)
    for left < right {
        mid := (left + right) / 2
        if arr[mid] <= target {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return left
}
func binarySearchGE(arr []int, target int) int {
    left := 0
    right := len(arr)
    for left < right {
        mid := (left + right) / 2
        if arr[mid] >= target {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return len(arr) - left
}