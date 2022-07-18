package struct_pattern

import "fmt"

type Type int

const (
	Common_Type = iota
	Special_Type
)

type Item struct {
	Id     int
	Name   string
	Length int
	Type   Type
}

type FilterAction interface {
	Filter(items []*Item)
}

type FilterContext struct {
	criterias []Criteria
}

func NewFilterContext(criterias []Criteria) *FilterContext {
	return &FilterContext{criterias: criterias}
}

func (f *FilterContext) filter(items []*Item) string {
	resItems := items
	result := ""
	for _, criteria := range f.criterias {
		resItems = criteria.Match(resItems)
	}
	for _, resItem := range resItems {
		result += fmt.Sprintf("%+v \n", resItem)
	}
	return result
}

type Criteria interface {
	Match(objs []*Item) []*Item
}

type NameCriteria struct {
	name string
}

func NewNameCriteria(name string) *NameCriteria {
	return &NameCriteria{name: name}
}

func (c *NameCriteria) Match(items []*Item) []*Item {
	list := make([]*Item, 0)
	for _, item := range items {
		if item.Name == c.name {
			list = append(list, item)
		}
	}
	return list
}

type LengthCriteria struct {
	minLength int
	maxLength int
}

func NewLengthCriteria(minLength int, maxLength int) *LengthCriteria {
	return &LengthCriteria{minLength: minLength, maxLength: maxLength}
}

func (c *LengthCriteria) Match(items []*Item) []*Item {
	list := make([]*Item, 0)
	for _, item := range items {
		if item.Length > c.minLength && item.Length < c.maxLength {
			list = append(list, item)
		}
	}
	return list
}

type TypeCriteria struct {
	typ Type
}

func NewTypeCriteria(typ Type) *TypeCriteria {
	return &TypeCriteria{typ: typ}
}

func (c *TypeCriteria) Match(items []*Item) []*Item {
	list := make([]*Item, 0)
	for _, item := range items {
		if item.Type == c.typ {
			list = append(list, item)
		}
	}
	return list
}
