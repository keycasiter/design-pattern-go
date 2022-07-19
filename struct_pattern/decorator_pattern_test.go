package struct_pattern

import (
	"fmt"
	"testing"
)

func TestDecoratorPattern(t *testing.T) {
	saltPie := NewSaltDecorator(NewPie())
	fmt.Println(saltPie.Cook())

	saltOilPie := NewSaltDecorator(NewOilDecorator(NewPie()))
	fmt.Println(saltOilPie.Cook())

	saltNoodles := NewSaltDecorator(NewNoodles())
	fmt.Println(saltNoodles.Cook())

	saltOilNoodles := NewSaltDecorator(NewOilDecorator(NewNoodles()))
	fmt.Println(saltOilNoodles.Cook())
}
