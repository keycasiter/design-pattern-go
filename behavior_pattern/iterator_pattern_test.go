package behavior_pattern

import (
	"fmt"
	"testing"
)

func TestIteratorPattern(t *testing.T) {
	items := []string{"1", "2", "3", "4", "5"}
	iterator := NewArrayIterator(items)

	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}
}
