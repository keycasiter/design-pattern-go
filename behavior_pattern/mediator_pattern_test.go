package behavior_pattern

import "testing"

func TestMediatorPattern(t *testing.T) {
	zs := NewUser("zhangsan")
	ls := NewUser("lisi")

	zs.Dail("Hi! lisi")
	ls.Dail("Hi! zhangsan")
}
