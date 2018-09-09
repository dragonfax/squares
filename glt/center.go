package glt

import "github.com/veandco/go-sdl2/sdl"

var _ ElementWidget = &Center{}
var _ Element = &CenterElement{}
var _ HasChild = &Center{}
var _ HasChildElement = &CenterElement{}

type Center struct {
	Child Widget
}

type CenterElement struct {
	widgetData
	sizeData
	parentData
	childElementData
}

func (ce *Center) getChild() Widget {
	return ce.Child
}

func (ce *Center) createElement() Element {
	element := &CenterElement{}
	element.widget = ce
	return element
}

func (ce *CenterElement) layout(c Constraints) error {

	ce.child.layout(c)

	ce.size = Size{width: c.maxWidth, height: c.maxHeight}

	childSize := ce.child.getSize()
	ce.child.getParentData().offset = Offset{
		x: (ce.size.width - childSize.width) / 2,
		y: (ce.size.height - childSize.height) / 2,
	}

	return nil
}

func (element *CenterElement) render(offset Offset, renderer *sdl.Renderer) {
	internalOffset := element.child.getParentData().offset
	offset.x += internalOffset.x
	offset.y += internalOffset.y
	element.child.render(offset, renderer)
}
