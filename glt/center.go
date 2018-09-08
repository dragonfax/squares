package glt

import "github.com/veandco/go-sdl2/sdl"

type Center struct {
	Child Widget
}

// test
var e Element = &CenterElement{}

type CenterElement struct {
	widget Widget
	child  Element
	sizeData
	parentData
}

func (ce *CenterElement) getChild() Widget {
	return ce.child
}

func (ce *CenterElement) layout(c Constraints) error {

	cw := ce.child.(Element)
	cw.layout(c)

	ce.size = Size{width: c.maxWidth, height: c.maxHeight}

	childSize := cw.getSize()

	cw.getParentData().offset = Offset{
		x: (ce.size.width - childSize.width) / 2,
		y: (ce.size.height - childSize.height) / 2,
	}

	return nil
}

func (c *CenterElement) render(offset Offset, renderer *sdl.Renderer) {
	cchild := c.child.(Element)
	internalOffset := cchild.getParentData().offset
	offset.x += internalOffset.x
	offset.y += internalOffset.y
	cchild.render(offset, renderer)
}
