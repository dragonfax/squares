package glt

/* A lazy way to do mixins, using embedded structs
 *
 * Not good enough for adding optional user parameters like Child to a literal struct,
 * as the user would have to type something weird like Struct{Child.Child: child}.
 * As a result we don't use this for Widgets, but will use it for elements.
 * */

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

type childElementData struct {
	child Element
}

func (ce childElementData) getChildElement() Element {
	return ce.child
}

func (ce *childElementData) setChildElement(child Element) {
	ce.child = child
}

type childrenElementsData struct {
	children []Element
}

func (ce childrenElementsData) getChildrenElements() []Element {
	return ce.children
}

func (ce *childrenElementsData) setChildrenElements(children []Element) {
	ce.children = children
}
