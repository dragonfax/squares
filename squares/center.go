package squares

import (
	"math"

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
		width:  constraintCenterDimension(c.maxWidth, childSize.width),
		height: constraintCenterDimension(c.maxHeight, childSize.height),
	})

	ce.child.setOffset(Offset{
		x: (ce.size.width - childSize.width) / 2,
		y: (ce.size.height - childSize.height) / 2,
	})

	return nil
}

// Golang needs the ternary operator
func constraintCenterDimension(constraint, child float64) float64 {
	if math.IsInf(constraint, 1) {
		return child
	}
	return math.Inf(1)
}

func (element *CenterElement) render(offset Offset, renderer *sdl.Renderer) {
	internalOffset := element.child.getOffset()
	offset.x += internalOffset.x
	offset.y += internalOffset.y
	element.child.render(offset, renderer)
}
