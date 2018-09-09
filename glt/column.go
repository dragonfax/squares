package glt

import "github.com/veandco/go-sdl2/sdl"

var _ ElementWidget = &Column{}
var _ Element = &ColumnElement{}
var _ HasChildren = &Column{}
var _ HasChildrenElements = &ColumnElement{}

type Column struct {
	Children []Widget
}

func (c *Column) getChildren() []Widget {
	return c.Children
}

func (c *Column) createElement() Element {
	return &ColumnElement{widget: c}
}

type ColumnElement struct {
	widget *Column
	sizeData
	parentData
	childrenElementsData
}

func (ce *ColumnElement) GetWidget() Widget {
	return ce.widget
}

func (ce *ColumnElement) layout(c Constraints) error {

	ce.size = Size{0, 0}

	numChildren := uint16(len(ce.children))

	for _, child := range ce.children {

		// TODO not sure about this.
		// might need to do them one at a time. and see whats left for the others.
		con := Constraints{
			minWidth:  c.maxWidth,
			minHeight: c.minHeight / numChildren,
			maxWidth:  c.maxWidth,
			maxHeight: c.maxHeight / numChildren,
		}

		child.layout(con)

		childSize := child.getSize()
		ce.size.width = MaxUint16(childSize.width, ce.size.width)
		offsetHeight := ce.size.height
		ce.size.height += childSize.height

		child.getParentData().offset = Offset{
			x: 0,
			y: offsetHeight,
		}

	}

	return nil
}

func (c *ColumnElement) render(offset Offset, renderer *sdl.Renderer) {

	for _, child := range c.children {
		child.render(offset, renderer)
		offset.y += child.getSize().height
	}
}
