package build_pattern

import (
	"context"
	"fmt"
	"testing"
	"time"
	"unsafe"
)

func TestGetInstanceByLazyModeWithBug(t *testing.T) {
	ctx := context.Background()
	cnt := 10

	for i := 0; i < cnt; i++ {
		go func() {
			instance := GetInstanceByLazyModeWithBug(ctx)
			fmt.Printf("%v\n", unsafe.Pointer(instance))
		}()
	}
	time.Sleep(5 * time.Second)
}

func TestGetInstanceByLazyModeWithFix(t *testing.T) {
	ctx := context.Background()
	cnt := 10

	for i := 0; i < cnt; i++ {
		go func() {
			instance := GetInstanceByLazyModeWithFix(ctx)
			fmt.Printf("%v\n", unsafe.Pointer(instance))
		}()
	}
	time.Sleep(5 * time.Second)
}

func TestGetInstanceByHungryMode(t *testing.T) {
	ctx := context.Background()
	cnt := 10

	for i := 0; i < cnt; i++ {
		go func() {
			instance := GetInstanceByHungryMode(ctx)
			fmt.Printf("%v\n", unsafe.Pointer(instance))
		}()
	}
	time.Sleep(5 * time.Second)
}
