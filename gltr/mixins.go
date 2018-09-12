package gltr

/* A lazy way to do mixins, using embedded structs
 *
 * Not good enough for adding optional user parameters like Child to a literal struct,
 * as the user would have to type something weird like Struct{Child.Child: child}.
 * As a result we don't use this for Widgets, but will use it for elements.
 * */
type elementData struct {
	widget Widget
	size   Size
	offset Offset
	parent Element
}

func (ce *elementData) getParentElement() Element {
	return ce.parent
}

func (ce *elementData) setParentElement(parent Element) {
	// Watchout, this doesn't update the parents list of children
	ce.parent = parent
}

func (ce *elementData) GetWidget() Widget {
	return ce.widget
}

func (ce *elementData) updateWidget(widget Widget) {
	ce.widget = widget
}

func (sd elementData) getSize() Size {
	return sd.size
}

func (ed *elementData) getOffset() Offset {
	return ed.offset
}

func (ed *elementData) setOffset(offset Offset) {
	ed.offset = offset
}

type childElementData struct {
	child Element
}

func (ce childElementData) getChildElement() Element {
	return ce.child
}

func (ce *childElementData) setChildElement(parent Element, child Element) {
	if ce.child != nil {
		ce.child.setParentElement(nil)
	}
	ce.child = child
	ce.child.setParentElement(parent)
}

type childrenElementsData struct {
	children []Element
}

func (ce childrenElementsData) getChildrenElements() []Element {
	return ce.children
}

func (ce *childrenElementsData) setChildrenElements(parent Element, children []Element) {
	for _, child := range ce.children {
		if child != nil {
			child.setParentElement(nil)
		}
	}
	ce.children = children
	for _, child := range children {
		child.setParentElement(parent)
	}
}
