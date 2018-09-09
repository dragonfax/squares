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
	/* StatlessWidget interface is included here only so that we get the same signature
	 * for the Build() method. If we add more methods to the StatelessWidget interface,
	 * we may not want them for State, so in that case remove this embedded interface.
	 */
	StatelessWidget
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
