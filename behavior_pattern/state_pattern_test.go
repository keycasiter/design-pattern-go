package behavior_pattern

import "testing"

func TestStatePattern(t *testing.T) {
	NewStateContext(NewStartState()).Run()
	NewStateContext(NewEndState()).Run()
}
