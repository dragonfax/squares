package glt

import "github.com/veandco/go-sdl2/sdl"

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

type element interface {
	layout(c constraints) error
	getParentData() *parentData
	getSize() Size
	render(Offset, *sdl.Renderer)
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
