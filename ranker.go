package ranker

import (
	"sort"
)

type Occurrence struct {
	Name       string
	Count      int
	AppearedAt int
}

type Ranker struct {
	Occurrences map[string]*Occurrence
	WindowSize  int

	clock int
}

func NewRanker(windowSize int) *Ranker {
	return &Ranker{
		Occurrences: make(map[string]*Occurrence, windowSize),
		WindowSize:  windowSize,
	}
}

func (r *Ranker) Add(s string) {
	r.clock++

	if occurrence, ok := r.Occurrences[s]; ok {
		occurrence.Count++
		return
	}

	if len(r.Occurrences) == r.WindowSize {
		r.prune()
	}

	r.Occurrences[s] = &Occurrence{
		Name:       s,
		Count:      1,
		AppearedAt: r.clock,
	}
}

func (r *Ranker) prune() {
	occs := r.occSlice()
	delete(r.Occurrences, occs[len(occs)-1].Name)
}

func (r *Ranker) occSlice() []*Occurrence {
	occs := []*Occurrence{}
	for _, occ := range r.Occurrences {

		//fmt.Println("append", occs, occ)
		occs = append(occs, occ)
	}
	sort.Sort(byTop(occs))
	return occs
}

func (r *Ranker) Top(n int) []string {
	occs := r.occSlice()

	top := []string{}
	for i := 0; i < n; i++ {
		occ := occs[i].Name
		//fmt.Println("append", top, occ)
		top = append(top, occ)
	}
	return top
}

type byTop []*Occurrence

func (s byTop) Len() int {
	return len(s)
}

func (s byTop) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byTop) Less(i, j int) bool {
	if s[j].Count < s[i].Count {
		return true
	}

	if s[j].Count == s[i].Count && s[j].AppearedAt < s[i].AppearedAt {
		return true
	}

	return false
}
