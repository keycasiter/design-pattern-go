package build_pattern

import (
	"fmt"
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	factory := &AbstractFactory{}

	fmt.Println(factory.GetSubject(Math).TellYouSubject())
	fmt.Println(factory.GetSubject(English).TellYouSubject())
	fmt.Println(factory.GetSubject(Chinese).TellYouSubject())

	fmt.Println(factory.GetLevel(Easy).TellYouLevel())
	fmt.Println(factory.GetLevel(Mid).TellYouLevel())
	fmt.Println(factory.GetLevel(Hard).TellYouLevel())
}
