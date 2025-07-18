type MaxHeap []int

func (h *MaxHeap) Push(x int) {
    *h = append(*h, x)
    h.up(len(*h) - 1)
}

func (h *MaxHeap) Pop() int {
    top := (*h)[0]
    last := len(*h) - 1
    (*h)[0] = (*h)[last]
    *h = (*h)[:last]
    h.down(0)
    return top
}

func (h *MaxHeap) up(i int) {
    for i > 0 {
        parent := (i - 1) / 2
        if (*h)[i] > (*h)[parent] {
            (*h)[i], (*h)[parent] = (*h)[parent], (*h)[i]
            i = parent
        } else {
            break
        }
    }
}

func (h *MaxHeap) down(i int) {
    n := len(*h)
    for {
        left := 2*i + 1
        right := 2*i + 2
        largest := i

        if left < n && (*h)[left] > (*h)[largest] {
            largest = left
        }
        if right < n && (*h)[right] > (*h)[largest] {
            largest = right
        } 
        if largest == i {
            break
        }

        (*h)[i], (*h)[largest] = (*h)[largest], (*h)[i]
        i = largest
    }
}
type MinHeap []int

func (h *MinHeap) Push(x int) {
    *h = append(*h, x)
    h.up(len(*h) - 1)
}

func (h *MinHeap) Pop() int {
    top := (*h)[0]
    last := len(*h) - 1
    (*h)[0] = (*h)[last]
    *h = (*h)[:last]
    h.down(0)
    return top
}

func (h *MinHeap) up(i int) {
    for i > 0 {
        parent := (i - 1) / 2
        if (*h)[i] < (*h)[parent] {
            (*h)[i], (*h)[parent] = (*h)[parent], (*h)[i]
            i = parent
        } else {
            break
        }
    }
}

func (h *MinHeap) down(i int) {
    n := len(*h)
    for  {
        left := 2*i + 1
        right := 2*i + 2 
        largest := i

        if left < n && (*h)[left] > (*h)[largest] {
            largest = left
        }
        if right < n && (*h)[right] > (*h)[largest] {
            largest = right
        }
        if largest == i {
            break
        }

        (*h)[i], (*h)[largest] = (*h)[largest], (*h)[i]
        i = largest
    }
}
func minimumDifference(nums []int) int64 {
    n3 := len(nums) 
    n := n3 / 3

    part1 := make([]int64, n+1)
    var leftSum int64 = 0
    ql := &MaxHeap{}

    for i := 0; i < n; i++ {
        leftSum += int64(nums[i])
        ql.Push(nums[i])
    }
    part1[0] = leftSum

    for i := n; i < 2*n; i++ {
        leftSum += int64(nums[i])
        ql.Push(nums[i])
        leftSum -= int64(ql.Pop())
        part1[i-n+1] = leftSum
    }

    rightSum := int64(0)
    qr := &MinHeap{}
    for i := 3*n - 1; i >= 2*n; i-- {
        rightSum += int64(nums[i])
        qr.Push(nums[i])
    }
    ans := part1[n] - rightSum

    for i := 2*n - 1; i >= n; i-- {
        rightSum += int64(nums[i])
        qr.Push(nums[i])
        rightSum -= int64(qr.Pop())
        diff := part1[i - n] - rightSum
        if diff < ans  {
            ans = diff
        }
    }
    return ans
}