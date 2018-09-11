package glt

import "github.com/veandco/go-sdl2/sdl"

var _ ElementWidget = &Padding{}
var _ Element = &PaddingElement{}
var _ HasChild = &Padding{}
var _ HasChildElement = &PaddingElement{}

type Padding struct {
	Padding EdgeInsets
	Child   Widget
}

func (p *Padding) createElement() Element {
	pe := &PaddingElement{}
	pe.widget = p
	return pe
}

type PaddingElement struct {
	elementData
	childElementData
}

func (element *PaddingElement) layout(c Constraints) error {

	widget := element.widget.(*Padding)

	innerConstraints := c.addMargins(widget.Padding)

	element.child.layout(innerConstraints)

	// multi child containers would read the sizes from the children, and position them accordingly.
	childSize := element.child.getSize()
	paddedSize := childSize.addMargin(widget.Padding)
	element.size = paddedSize

	// offset for Padding is easy, just offset by the padding amount.
	element.child.setOffset(Offset{widget.Padding.All, widget.Padding.All})

	return nil
}

func (element *PaddingElement) render(offset Offset, renderer *sdl.Renderer) {
	internalOffset := element.child.getOffset()
	offset.x += internalOffset.x
	offset.y += internalOffset.y
	element.child.render(offset, renderer)
}
