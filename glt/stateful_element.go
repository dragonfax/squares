package glt

import "github.com/veandco/go-sdl2/sdl"

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
