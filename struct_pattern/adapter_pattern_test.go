package struct_pattern

import (
	"fmt"
	"testing"
)

func TestAdapterPattern(t *testing.T) {
	fmt.Println(NewEngineAdapter(Fast_Engine).Run())
	fmt.Println(NewEngineAdapter(Slow_Engine).Run())
}
