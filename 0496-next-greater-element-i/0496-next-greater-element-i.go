func nextGreaterElement(nums1 []int, nums2 []int) []int {
    stack := []int{}

    for i := 0; i < len(nums1); i++ {
        j := slices.Index(nums2, nums1[i])
        for j < len(nums2) {
            if nums2[j] > nums1[i] {
                stack = append(stack, nums2[j])
                break
            } 
            j++
        }
        if len(stack) == i {
            stack = append(stack, -1)
        }
    }
    return stack
}