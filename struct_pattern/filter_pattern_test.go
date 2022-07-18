package struct_pattern

import (
	"fmt"
	"testing"
)

var zhangsan10To20CommonType = []Criteria{
	NewNameCriteria("zhangsan"),
	NewLengthCriteria(10, 20),
	NewTypeCriteria(Common_Type),
}

var lisi20To100SpecialType = []Criteria{
	NewNameCriteria("lisi"),
	NewLengthCriteria(20, 100),
	NewTypeCriteria(Special_Type),
}

func TestFilterPattern(t *testing.T) {
	filterCtx := NewFilterContext(zhangsan10To20CommonType)
	fmt.Println("Result:", filterCtx.filter([]*Item{
		{
			Id:     1,
			Name:   "zhangsan",
			Length: 15,
			Type:   Common_Type,
		},
		{
			Id:     2,
			Name:   "lisi",
			Length: 55,
			Type:   Special_Type,
		},
	}))

	filterCtx2 := NewFilterContext(lisi20To100SpecialType)
	fmt.Println("Result:", filterCtx2.filter([]*Item{
		{
			Id:     1,
			Name:   "zhangsan",
			Length: 15,
			Type:   Common_Type,
		},
		{
			Id:     2,
			Name:   "lisi",
			Length: 55,
			Type:   Special_Type,
		},
	}))
}
