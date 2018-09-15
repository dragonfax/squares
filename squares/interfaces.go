package squares

import "github.com/veandco/go-sdl2/sdl"

type Widget interface {
}

type SetStateFunc func()

type BuildContext interface {
	GetWidget() Widget
}

type StatelessContext interface {
	// not sure what goes here yet.
}

type StatefulContext interface {
	SetState(SetStateFunc)
	GetWidget() Widget
}

type StatelessWidget interface {
	// Can't embed function types as interface methods, so we'll just copy the definition of BuildFunc
	Build(context StatelessContext) (Widget, error)
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
	Build(context StatefulContext) (Widget, error)
}

type Element interface {
	layout(c Constraints) error
	getOffset() Offset
	setOffset(Offset)
	getSize() Size
	render(Offset, *sdl.Renderer)
	updateWidget(Widget)
	BuildContext
	getParentElement() Element
	setParentElement(Element)
}

type HasChildElement interface {
	setChildElement(Element, Element)
	getChildElement() Element
}

type HasChildrenElements interface {
	getChildrenElements() []Element
	setChildrenElements(Element, []Element)
}

/* A widget that has a special Element just for it
 *	Such a widget won't have a Build() method,
 *	And may or may not have chilcdren
 */
type ElementWidget interface {
	createElement() Element
}
