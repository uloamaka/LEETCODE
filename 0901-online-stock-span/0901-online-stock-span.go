type Pair struct {
    price int
    span int
}

type StockSpanner struct {
    stack []Pair
}


func Constructor() StockSpanner {
    return StockSpanner{
        stack: []Pair{},
    }
}


func (this *StockSpanner) Next(price int) int {
    span := 1

    // pop all price <= current price, increamenting the span
    for len(this.stack) > 0 && this.stack[len(this.stack)-1].price <= price {
        span += this.stack[len(this.stack)-1].span
        this.stack = this.stack[:len(this.stack)-1]
    }

    // Push the current price with its span
    this.stack = append(this.stack, Pair{
        price: price,
        span: span,
    })

    return span
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */