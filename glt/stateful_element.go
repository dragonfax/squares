package glt

import "github.com/veandco/go-sdl2/sdl"

type StatefulElement struct {
	widget Widget
	state  State
	child  Element
}

var _ Element = StatefulElement{}

func (se StatefulElement) getSize() Size {
	return se.child.getSize()
}

func (se StatefulElement) getParentData() *parentData {
	return se.child.getParentData()
}

func (se StatefulElement) layout(c Constraints) error {
	return se.child.layout(c)
}

func (se StatefulElement) render(o Offset, r *sdl.Renderer) {
	se.child.render(o, r)
}
