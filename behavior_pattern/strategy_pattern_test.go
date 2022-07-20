package behavior_pattern

import (
	"fmt"
	"testing"
)

func TestStrategyPattern(t *testing.T) {
	fmt.Println(NewStrategyContext(NewSumStrategy()).Run(1, 2))
	fmt.Println(NewStrategyContext(NewMultiplyStrategy()).Run(1, 2))
}
