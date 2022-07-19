package behavior_pattern

import "fmt"

type Processor struct {
	name string
	next *Processor
}

func NewProcessor(name string, next *Processor) *Processor {
	return &Processor{name: name, next: next}
}

func (p *Processor) Process() {
	fmt.Println("[process]-> ", p.name)
	if nil != p.next {
		p.next.Process()
	}
}
