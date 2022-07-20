package behavior_pattern

import (
	"fmt"
	"testing"
)

func TestMementoPattern(t *testing.T) {
	sm := NewStateMachine()
	sm.Add(NewMemento("#1"))
	sm.Add(NewMemento("#2"))
	sm.Add(NewMemento("#3"))
	sm.Add(NewMemento("#4"))

	fmt.Println(sm.GetMemento("#1"))
	fmt.Println(sm.GetMemento("#2"))
	fmt.Println(sm.GetMemento("#3"))
}
