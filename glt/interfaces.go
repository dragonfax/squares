package glt

import "github.com/veandco/go-sdl2/sdl"

type Widget interface {
}

type StatelessWidget interface {
	Build() (Widget, error)
}

type HasChild interface {
	getChild() Widget
}

type HasChildren interface {
	getChildren() []Widget
}

type StatefulWidget interface {
	CreateState() State
}

type State interface {
	Build() Widget
}

type Element interface {
	layout(c Constraints) error
	getParentData() *parentData
	getSize() Size
	render(Offset, *sdl.Renderer)
}

type HasChildElement interface {
	setChildElement(Element)
}

type HasChildrenElements interface {
	setChildrenElements([]Element)
}

/* A widget that has a special Element just for it
 *	Such a widget won't have a Build() method,
 *	And may or may not have chilcdren
 */
type ElementWidget interface {
	createElement() Element
}
