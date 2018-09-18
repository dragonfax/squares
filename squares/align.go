package squares

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

var _ ElementWidget = Align{}
var _ Element = &AlignElement{}
var _ HasChild = Align{}
var _ HasChildElement = &AlignElement{}

type AlignmentGeometry uint8

const (
	AlignmentNone AlignmentGeometry = iota
	AlignmentBottomLeft
	AlignmentUpperRight
)

type Align struct {
	Alignment AlignmentGeometry
	Child     Widget
}

func (a Align) getChild() Widget {
	return a.Child
}

func (a Align) createElement() Element {
	ae := &AlignElement{}
	ae.widget = a
	return ae
}

type AlignElement struct {
	elementData
	childElementData
}

func (ae *AlignElement) layout(c Constraints) error {

	// child size
	ae.child.layout(c)
	childSize := ae.child.getSize()

	// our size
	size := Size{Width: -1, Height: -1}
	if math.IsInf(c.maxHeight, 1) {
		// unbound height, use childs
		size.Height = childSize.Height
	}
	if math.IsInf(c.maxWidth, 1) {
		size.Width = childSize.Width
	}
	ae.size = c.constrain(size)

	// child offset
	switch ae.widget.(Align).Alignment {
	case AlignmentBottomLeft:
		x := 0.0
		y := ae.size.Height - childSize.Height
		ae.child.setOffset(Offset{x: x, y: y})
	case AlignmentUpperRight:
		y := 0.0
		x := ae.size.Width - childSize.Width
		ae.child.setOffset(Offset{x: x, y: y})
	default:
		panic("unimplemented")
	}

	return nil
}

func (ae *AlignElement) render(o Offset, r *sdl.Renderer) {
	ae.child.render(o.Add(ae.child.getOffset()), r)
}
