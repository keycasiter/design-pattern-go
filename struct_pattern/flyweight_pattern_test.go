package struct_pattern

import (
	"fmt"
	"testing"
)

func TestFlyweightPattern(t *testing.T) {
	fmt.Println(NewCarFactory().GetCar(Yellow).Drive())
	fmt.Println(NewCarFactory().GetCar(Red).Drive())
	fmt.Println(NewCarFactory().GetCar(Green).Drive())
}
