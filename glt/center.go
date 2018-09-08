package flutter

import "github.com/veandco/go-sdl2/sdl"

type Center struct {
	Child Widget
	sizeData
	parentData
}

func (ce *Center) getChild() Widget {
	return ce.Child
}

func (ce *Center) setChild(c Widget) {
	ce.Child = c
}

func (ce *Center) layout(c constraints) error {

	cw := ce.Child.(coreWidget)
	cw.layout(c)

	ce.size = Size{width: c.maxWidth, height: c.maxHeight}

	childSize := cw.getSize()

	cw.getParentData().offset = Offset{
		x: (ce.size.width - childSize.width) / 2,
		y: (ce.size.height - childSize.height) / 2,
	}

	return nil
}

func (c *Center) render(offset Offset, renderer *sdl.Renderer) {
	cchild := c.Child.(coreWidget)
	internalOffset := cchild.getParentData().offset
	offset.x += internalOffset.x
	offset.y += internalOffset.y
	cchild.render(offset, renderer)
}
