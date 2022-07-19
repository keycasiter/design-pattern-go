package struct_pattern

import "fmt"

type Food interface {
	Cook() string
}

type Noodles struct {
	Material string
}

func NewNoodles() *Noodles {
	return &Noodles{Material: "noodles"}
}

func (f *Noodles) Cook() string {
	return f.Material
}

type Pie struct {
	Material string
}

func NewPie() *Pie {
	return &Pie{Material: "pie"}
}

func (f *Pie) Cook() string {
	return f.Material
}

type SaltDecorator struct {
	Material string
	Food     Food
}

func NewSaltDecorator(food Food) *SaltDecorator {
	return &SaltDecorator{Material: "salt", Food: food}
}

func (f *SaltDecorator) Cook() string {
	return fmt.Sprintf("%s add:%s ", f.Food.Cook(), f.Material)
}

type OilDecorator struct {
	Material string
	Food     Food
}

func NewOilDecorator(food Food) *OilDecorator {
	return &OilDecorator{Material: "oil", Food: food}
}

func (f *OilDecorator) Cook() string {
	return fmt.Sprintf("%s add:%s ", f.Food.Cook(), f.Material)
}
