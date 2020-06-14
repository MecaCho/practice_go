package weekly_192

type BrowserHistory struct {
	Pages []string
	Loc   int
}

func Constructor(homepage string) BrowserHistory {

	pages := []string{homepage}

	return BrowserHistory{pages, 0}

}

func (this *BrowserHistory) Visit(url string) {
	if this.Loc == len(this.Pages)-1 {
		this.Pages = append(this.Pages, url)
	} else {
		this.Pages = this.Pages[:this.Loc+1]
		this.Pages = append(this.Pages, url)
	}
	this.Loc++
	// fmt.Println(this.Pages)
}

func (this *BrowserHistory) Back(steps int) string {
	if this.Loc-steps < 0 {
		this.Loc = 0
	} else {
		this.Loc -= steps
	}
	return this.Pages[this.Loc]
}

func (this *BrowserHistory) Forward(steps int) string {
	if this.Loc+steps > len(this.Pages)-1 {
		this.Loc = len(this.Pages) - 1
	} else {
		this.Loc += steps
	}
	return this.Pages[this.Loc]
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
