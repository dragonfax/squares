package squares

import (
	"github.com/veandco/go-sdl2/sdl"
)

var _ ElementWidget = &Center{}
var _ Element = &CenterElement{}
var _ HasChild = &Center{}
var _ HasChildElement = &CenterElement{}

type Center struct {
	Child Widget
}

type CenterElement struct {
	elementData
	childElementData
}

func (ce *Center) createElement() Element {
	element := &CenterElement{}
	element.widget = ce
	return element
}

func (ce *CenterElement) layout(c Constraints) error {

	if ce.child == nil {
		panic("center with no child element")
	}

	ce.child.layout(c.loosen())

	childSize := ce.child.getSize()
	ce.size = c.constrain(Size{
		Width:  constraintCenterDimension(c.maxWidth, childSize.Width),
		Height: constraintCenterDimension(c.maxHeight, childSize.Height),
	})

	ce.child.setOffset(Offset{
		x: (ce.size.Width - childSize.Width) / 2,
		y: (ce.size.Height - childSize.Height) / 2,
	})

	return nil
}

// Golang needs the ternary operator
func constraintCenterDimension(constraint, child float64) float64 {
	if constraint == Inf {
		return child
	}
	return Inf
}

func (element *CenterElement) render(offset Offset, renderer *sdl.Renderer) {
	internalOffset := element.child.getOffset()
	offset.x += internalOffset.x
	offset.y += internalOffset.y
	element.child.render(offset, renderer)
}
