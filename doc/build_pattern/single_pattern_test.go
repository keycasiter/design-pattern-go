package build_pattern

import (
	"context"
	"fmt"
	"testing"
)

func TestGetInstanceByLazyModeWithBug(t *testing.T) {
	ctx := context.Background()
	cnt := 10
	holder := make([]*Instance, 0)

	for i := 0; i < cnt; i++ {
		instance := GetInstanceByLazyModeWithBug(ctx)
		holder = append(holder, instance)
	}
	for i := 0; i < cnt; i++ {
		fmt.Printf("%v\n", &holder[i])
	}
}

func TestGetInstanceByHungryMode(t *testing.T) {

}
