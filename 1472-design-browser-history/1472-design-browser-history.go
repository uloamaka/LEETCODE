type BrowserHistory struct {
    window     []string
    currentTab int
}

func Constructor(homepage string) BrowserHistory {
    return BrowserHistory{
        window:     []string{homepage},
        currentTab: 0,
    }
}

func (bh *BrowserHistory) Visit(url string) {
    bh.window = bh.window[:bh.currentTab+1]

    bh.window = append(bh.window, url)
    bh.currentTab++
}

func (bh *BrowserHistory) Back(steps int) string {
    if steps > bh.currentTab {
        bh.currentTab = 0
    } else {
        bh.currentTab -= steps
    }
    return bh.window[bh.currentTab]
}

func (bh *BrowserHistory) Forward(steps int) string {
    maxForward := len(bh.window) - 1 - bh.currentTab
    if steps > maxForward {
        bh.currentTab += maxForward
    } else {
        bh.currentTab += steps
    }
    return bh.window[bh.currentTab]
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */