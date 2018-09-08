package glt

import "github.com/veandco/go-sdl2/sdl"

type Center struct {
	Child Widget
}

// test
var e element = &centerElement{}

type centerElement struct {
	widget Widget
	child  element
	sizeData
	parentData
}

func (ce *centerElement) getChild() Widget {
	return ce.child
}

func (ce *centerElement) layout(c constraints) error {

	cw := ce.child.(element)
	cw.layout(c)

	ce.size = Size{width: c.maxWidth, height: c.maxHeight}

	childSize := cw.getSize()

	cw.getParentData().offset = Offset{
		x: (ce.size.width - childSize.width) / 2,
		y: (ce.size.height - childSize.height) / 2,
	}

	return nil
}

func (c *centerElement) render(offset Offset, renderer *sdl.Renderer) {
	cchild := c.child.(element)
	internalOffset := cchild.getParentData().offset
	offset.x += internalOffset.x
	offset.y += internalOffset.y
	cchild.render(offset, renderer)
}
