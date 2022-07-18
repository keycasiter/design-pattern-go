package build_pattern

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	builder := NewConfigBuilder()
	builder.WithHost("127.0.0.1")
	builder.WithPort(3306)
	builder.WithMaxNum(10)
	builder.Build()

	fmt.Printf("builder:%+v", builder)
}
