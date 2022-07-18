package build_pattern

import (
	"fmt"
	"testing"
)

func TestFactoryPattern(t *testing.T) {
	factory := &CarFactory{}
	teslaProducer := factory.GetProducer(Tesla)
	fmt.Println(teslaProducer.Produce())
	fordProducer := factory.GetProducer(Ford)
	fmt.Println(fordProducer.Produce())
	toytoProducer := factory.GetProducer(Toyto)
	fmt.Println(toytoProducer.Produce())
}
