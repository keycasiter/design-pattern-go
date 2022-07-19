package struct_pattern

import (
	"fmt"
	"testing"
)

func TestFacadePattern(t *testing.T) {
	fmt.Println(NewFacadeSrv().Query())
	fmt.Println(NewFacadeSrv().Create())
	fmt.Println(NewFacadeSrv().Update())
}
