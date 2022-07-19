package struct_pattern

import (
	"fmt"
	"testing"
)

func TestProxyPattern(t *testing.T) {
	fmt.Println(NewSrv().Do())

	fmt.Println(NewProxyLogSrv().Do())
}
