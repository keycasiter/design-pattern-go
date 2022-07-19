package struct_pattern

type Color int

const (
	Red = iota
	Yellow
	Green
)

type Car interface {
	Drive() string
}

type YellowCar struct {
	color string
}

func NewYellowCar() *YellowCar {
	return &YellowCar{color: "yellow"}
}

func (c *YellowCar) Drive() string {
	return "yellow car drive"
}

type RedCar struct {
	color string
}

func NewRedCar() *RedCar {
	return &RedCar{color: "red"}
}

func (c *RedCar) Drive() string {
	return "red car drive"
}

type GreenCar struct {
	color string
}

func NewGreenCar() *GreenCar {
	return &GreenCar{color: "green"}
}

func (c *GreenCar) Drive() string {
	return "green car drive"
}

type CarFactory struct {
	colorCarMap map[Color]Car
}

func NewCarFactory() *CarFactory {
	colorCarMap := make(map[Color]Car, 0)
	colorCarMap[Yellow] = NewYellowCar()
	colorCarMap[Red] = NewRedCar()
	colorCarMap[Green] = NewGreenCar()
	return &CarFactory{colorCarMap: colorCarMap}
}

func (f *CarFactory) GetCar(color Color) Car {
	if car, ok := f.colorCarMap[color]; ok {
		return car
	}
	return nil
}
