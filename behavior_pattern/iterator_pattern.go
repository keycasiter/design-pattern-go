package behavior_pattern

type Iterator interface {
	HasNext() bool
	Next() string
}

type ArrayIterator struct {
	items []string
	index int
}

func NewArrayIterator(items []string) *ArrayIterator {
	return &ArrayIterator{items: items}
}

func (a *ArrayIterator) HasNext() bool {
	if len(a.items) == a.index {
		return false
	}
	return true
}

func (a *ArrayIterator) Next() string {
	if !a.HasNext() {
		return ""
	}
	a.index++
	return a.items[a.index-1]
}
