package flutter

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type BuildContext struct {
}

type Widget interface {
}

type hasBuild interface { // StatelessWidget
	Build(*BuildContext) (Widget, error)
}

type hasChild interface { // Container
	getChild() Widget
	setChild(Widget)
}

type hasChildren interface {
	getChildren() []Widget
	setChildren([]Widget)
}

type hasRender interface { // RenderObject
	Render(*sdl.Renderer) error
}

type EdgeInsets struct {
	All uint16
}

type Offset struct {
	x, y uint16
}

type Size struct {
	width, height uint16
}

type sizeData struct {
	size Size
}

func (sd sizeData) getSize() Size {
	return sd.size
}

type parentData struct {
	offset Offset
}

// let an embedded struct return itself in order to match interfaces (interfaces for struct elements)
func (ce parentData) getParentData() *parentData {
	return &ce
}

type coreWidget interface {
	layout(c constraints) error
	getParentData() *parentData
	getSize() Size
	render(Offset, *sdl.Renderer)
}

// use MaxUint32 for +Inf during layout
type constraints struct {
	minWidth, minHeight, maxWidth, maxHeight uint16
}

func (s Size) addMargin(in EdgeInsets) Size {
	return Size{width: s.width + in.All, height: s.height + in.All}
}

func (c constraints) addMargins(in EdgeInsets) constraints {
	// TODO fix the math here
	if c.minWidth > in.All {
		c.minWidth -= in.All
	} else {
		c.minWidth = 0
	}

	if c.minHeight > in.All {
		c.minHeight -= in.All
	} else {
		c.minHeight = 0
	}

	if c.maxWidth != math.MaxUint16 {
		c.maxWidth -= in.All
	}
	if c.maxHeight != math.MaxUint16 {
		c.maxHeight -= in.All
	}
	return c
}

func MaxUint16(a, b uint16) uint16 {
	if a > b {
		return a
	} else {
		return b
	}
}
