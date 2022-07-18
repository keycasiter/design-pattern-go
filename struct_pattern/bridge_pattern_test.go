package struct_pattern

import (
	"fmt"
	"testing"
)

func TestDraw(t *testing.T) {
	fmt.Println(NewDraw(&CircleShape{}, &RedColor{}).Draw())
	fmt.Println(NewDraw(&CircleShape{}, &BlueColor{}).Draw())
	fmt.Println(NewDraw(&StarShape{}, &RedColor{}).Draw())
	fmt.Println(NewDraw(&StarShape{}, &BlueColor{}).Draw())
}
