package glt

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type Widget interface {
}

type statelessWidget interface {
	Build() (Widget, error)
}

type hasChild interface {
	getChild() Widget
}

type hasChildren interface {
	getChildren() []Widget
}

type statefulWidget interface {
	CreateState() State
}

type State interface {
	Build() Widget
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

type element interface {
	layout(c constraints) error
	getParentData() *parentData
	getSize() Size
	render(Offset, *sdl.Renderer)
}

type statefulElement struct {
	widget Widget
	state  State
	child  element
}

var _ element = statefulElement{}

func (se statefulElement) getSize() Size {
	return se.child.getSize()
}

func (se statefulElement) getParentData() *parentData {
	return se.child.getParentData()
}

func (se statefulElement) layout(c constraints) error {
	return se.child.layout(c)
}

func (se statefulElement) render(o Offset, r *sdl.Renderer) {
	se.child.render(o, r)
}

type hasChildElement interface {
	setChild(element)
}

type hasChildrenElements interface {
	setChildrenElements([]element)
}

/* A widget that has a special Element just for it
 *	Such a widget won't have a Build() method,
 *	And may or may not have chilcdren
 */
type elementWidget interface {
	createElement() element
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
