package flutter

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type BuildContext struct {
}

type Widget interface {
}

type HasBuild interface { // StatelessWidget
	Build(*BuildContext) (Widget, error)
}

type HasChild interface { // Container
	GetChild() Widget
	SetChild(Widget)
}

type HasChildren interface {
	GetChildren() []Widget
	SetChildren([]Widget)
}

type HasRender interface { // RenderObject
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

type ParentData struct {
	offset Offset
}

type CoreWidget interface {
	layout(c Constraints) error
	getParentData() *ParentData
	getSize() Size
}

// use MaxUint32 for +Inf during layout
type Constraints struct {
	minWidth, minHeight, maxWidth, maxHeight uint16
}

func (s Size) addMargin(in EdgeInsets) Size {
	return Size{width: s.width + in.All, height: s.height + in.All}
}

func (c Constraints) addMargins(in EdgeInsets) Constraints {
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
