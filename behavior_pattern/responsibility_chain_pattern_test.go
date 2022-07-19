package behavior_pattern

import "testing"

func TestResponsibilityChainPattern(t *testing.T) {
	p4 := NewProcessor("4", nil)
	p3 := NewProcessor("3", p4)
	p2 := NewProcessor("2", p3)
	p1 := NewProcessor("1", p2)
	p1.Process()
}
