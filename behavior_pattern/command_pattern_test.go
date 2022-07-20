package behavior_pattern

import (
	"fmt"
	"testing"
)

func TestCommandPattern(t *testing.T) {
	fmt.Println(NewExecutor([]Command{
		NewDoCommand(),
		NewQueryCommand(),
	}).Run())
}
