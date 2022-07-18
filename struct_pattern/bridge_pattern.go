package struct_pattern

import "fmt"

type DrawAction interface {
	Draw() string
}

type Draw struct {
	shapeAction ShapeAction
	colorAction ColorAction
}

func NewDraw(shapeAction ShapeAction, colorAction ColorAction) *Draw {
	return &Draw{shapeAction: shapeAction, colorAction: colorAction}
}

func (d *Draw) Draw() string {
	return fmt.Sprintf("color:%s , shape:%s", d.colorAction.Color(), d.shapeAction.Shape())
}

//定义形状
type ShapeAction interface {
	Shape() string
}

type CircleShape struct {
}

func (shape *CircleShape) Shape() string {
	return "Circle"
}

type StarShape struct {
}

func (shape *StarShape) Shape() string {
	return "Star"
}

//定义颜色
type ColorAction interface {
	Color() string
}

type RedColor struct {
}

func (color *RedColor) Color() string {
	return "Red"
}

type BlueColor struct {
}

func (color *BlueColor) Color() string {
	return "Blue"
}
