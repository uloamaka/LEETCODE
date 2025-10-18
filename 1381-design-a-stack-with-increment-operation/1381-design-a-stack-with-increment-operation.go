type CustomStack struct {
    stack []int
    maxSize int
}


func Constructor(maxSize int) CustomStack {
    return CustomStack{ 
        stack: make([]int, 0, maxSize), 
        maxSize: maxSize,
    }
}


func (this *CustomStack) Push(x int)  {
    if len(this.stack) < this.maxSize {
        this.stack = append(this.stack, x)
    }  
}


func (this *CustomStack) Pop() int {
    n := len(this.stack)
    if n ==  0 {
        return -1
    }
    elem := this.stack[n-1]
    this.stack = this.stack[:n-1]
    return elem
}


func (this *CustomStack) Increment(k int, val int)  {
    m := this.maxSize
    for i := 0; i < k; i++ {
        if m > 0 {
            this.stack[i] += val
            m--
        } else {
            break
        }
    }
}


/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */