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
	return &PaddingElement{widget: p}
}

func (p *Padding) getChild() Widget {
	return p.Child
}

func (p *Padding) setChild(child Widget) {
	p.Child = child
}

type PaddingElement struct {
	widget *Padding
	sizeData
	parentData
	childElementData
}

func (pe *PaddingElement) getWidget() Widget {
	return pe.widget
}

func (element *PaddingElement) layout(c Constraints) error {

	innerConstraints := c.addMargins(element.widget.Padding)

	element.child.layout(innerConstraints)

	// multi child containers would read the sizes from the children, and position them accordingly.
	childSize := element.child.getSize()
	paddedSize := childSize.addMargin(element.widget.Padding)
	element.size = paddedSize

	// offset for Padding is easy, just offset by the padding amount.
	element.child.getParentData().offset = Offset{element.widget.Padding.All, element.widget.Padding.All}

	return nil
}

func (element *PaddingElement) render(offset Offset, renderer *sdl.Renderer) {
	internalOffset := element.child.getParentData().offset
	offset.x += internalOffset.x
	offset.y += internalOffset.y
	element.child.render(offset, renderer)
}
