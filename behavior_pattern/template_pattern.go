package behavior_pattern

import "fmt"

type Template interface {
	Do()
	Init()
	Process()
	End()
}

type BaseTemplate struct {
	Template
}

func NewBaseTemplate(template Template) *BaseTemplate {
	return &BaseTemplate{Template: template}
}

func (base BaseTemplate) Do() {
	base.Init()

	base.Process()

	base.End()
}

type Fast struct {
	name string
	BaseTemplate
}

func NewFast() *Fast {
	return &Fast{name: "fast"}
}

func (f *Fast) Init() {
	fmt.Printf("%s init\n", f.name)
}

func (f *Fast) Process() {
	fmt.Printf("%s process\n", f.name)
}

func (f *Fast) End() {
	fmt.Printf("%s end\n", f.name)
}

type Slow struct {
	name string
	BaseTemplate
}

func NewSlow() *Slow {
	return &Slow{name: "slow"}
}

func (f *Slow) Init() {
	fmt.Printf("%s init\n", f.name)
}

func (f *Slow) Process() {
	fmt.Printf("%s process\n", f.name)
}

func (f *Slow) End() {
	fmt.Printf("%s end\n", f.name)
}
