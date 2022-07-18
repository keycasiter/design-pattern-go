package build_pattern

import (
	"fmt"
	"testing"
)

func TestOptionBuilder(t *testing.T) {
	cfg := NewConfig(
		WithHost("127.0.0.1"),
		WithPort(80),
		WithMaxNum(10),
	)
	fmt.Printf("cfg:%+v", cfg)
}
