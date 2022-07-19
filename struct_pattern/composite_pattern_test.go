package struct_pattern

import "testing"

func TestCompositePattern(t *testing.T) {
	level_1 := NewNode("level-1")
	level_2_1 := NewNode("level-2-1")
	level_2_2 := NewNode("level-2-2")

	level_1.AddChildNodes([]*Node{level_2_1, level_2_2})

	level_1.PrintLevel()
}
