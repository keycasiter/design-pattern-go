package build_pattern

type ProductMode int

const (
	Toyto = iota
	Tesla
	Ford
)

type CarFactory struct{}

type CarProducer interface {
	Produce() string
}

//Toyto
type ToytoProducer struct{}

func (p *ToytoProducer) Produce() string {
	return "Toyto Car"
}

//Tesla
type TeslaProducer struct{}

func (p *TeslaProducer) Produce() string {
	return "Tesla Car"
}

//Ford
type FordProducer struct{}

func (p *FordProducer) Produce() string {
	return "Ford Car"
}

func (f *CarFactory) GetProducer(productMode ProductMode) CarProducer {
	switch productMode {
	case Toyto:
		return &ToytoProducer{}
	case Tesla:
		return &TeslaProducer{}
	case Ford:
		return &FordProducer{}
	default:
		panic("unsupport productMode")
	}
}
